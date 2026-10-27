package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int64; F1 int8; F2 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 uint64; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32 }
func genFn0(p0 uint16) string {
    p0 += 2
    if (42 >= 1) {
        if false {
            for fi6 := 0; fi6 < 4; fi6++ {
                v7 := GenS1{F0: 2147483647, F1: (("\n\t" + "hello") + "\n\t"), F2: 1}
                fmt.Println(p0)
                _ = fi6
                _ = v7
            }
        } else {
            var v8 uint32 = 129
            fmt.Println(v8, p0)
            var v9 string = "\n\t"
            _ = v8
            _ = v9
        }
        fmt.Println(p0)
    }
    var v10 bool = true
    v11 := GenS1{F0: -766, F1: "", F2: 1}
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 int8 = -41
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS2{F0: 222, F1: -1}
            fmt.Println(fi12)
            _ = fi15
            _ = v16
        }
        _ = fi12
    }
    _ = p0
    _ = v10
    _ = v11
    return ("\n\t" + "\n\t")
}
func genFn1() uint16 {
    v17 := GenS3{F0: 42, F1: 1, F2: 9223372036854775807}
    switch "" {
    case "":
        var v18 uint16 = 2
        _ = v18
    case "a":
        if true {
            v19 := []int{403, 694, 4294967295}
            var v20 bool = ((!true) && (127 < 0))
            v20 = false
            _ = v19
            _ = v20
        }
    case "hello":
        if true {
            if (true || true) {
                for fi21 := 0; fi21 < 2; fi21++ {
                    fmt.Println(fi21)
                    var v22 uint16 = 32767
                    v22 -= v22
                    _ = fi21
                    _ = v22
                }
                var v23 string = "a"
                v23 = v23
                _ = v23
            } else {
                var v24 int64 = 42
                v25 := GenS2{F0: 4294967295, F1: v24}
                for fi26 := 0; fi26 < 10; fi26++ {
                    var v27 uint8 = 179
                    v28 := GenS4{F0: (("a" + "\n\t") + "foo bar")}
                    v24 += v24
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v24
                _ = v25
            }
            v29 := []int{}
            var v30 uint8 = 0
            _ = v29
            _ = v30
        } else {
            v31 := GenS5{F0: (42 <= 128), F1: 1, F2: 0}
            var v32 uint8 = 127
            var v33 int64 = 408
            _ = v31
            _ = v32
            _ = v33
        }
    default:
        switch ("hello" + "\n\t") {
        case "foo bar":
            v34 := []int{-814, 565}
            _ = v34
        default:
            var v35 string = "a"
            _ = v35
        }
    }
    var v36 uint64 = 2
    _ = v17
    _ = v36
    return 255
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint16 = 1
    v3 := []uint32{4294967295, 2147483648}
    var v4 int16 = -676
    v5 := make(map[string]bool)
    var v37 bool = (true && false)
    if true {
        var v38 float32 = 0.1
        delete(v5, (("hello" + "a") + ""))
        v39 := GenS5{F0: (1 < 255), F1: (^(v1)), F2: 272}
        _ = v38
        _ = v39
    }
    var v40 float32 = (3.14159 * 0.1)
    for fi41 := 0; fi41 < 15; fi41++ {
        for fi42 := 0; fi42 < 12; fi42++ {
            var v43 bool = (0 >= -1)
            _ = len(v5)
            v1 = (v1 + v1)
            _ = fi42
            _ = v43
        }
        v44 := GenS6{F0: -1}
        _ = fi41
        _ = v44
    }
    v45 := GenS7{F0: -384}
    delete(v5, ("a" + ""))
    v1 += v1
    if false {
        v46 := GenS4{F0: ""}
        _ = v46
    }
    var v47 uint = 18446744073709551615
    for fi48 := 0; fi48 < 18; fi48++ {
        if true {
            v40 += (v40 + float32(((1.5 / (1.5 / 1.0)) / 1e10)))
            var v49 uint16 = (v2 >> ((749) & 63))
            v50 := GenS8{F0: -2147483648}
            _ = v49
            _ = v50
        } else {
            var v51 uint16 = uint16(v1)
            _ = v51
        }
        var v52 bool = true
        for fi53 := 0; fi53 < 19; fi53++ {
            v4 = (v4 ^ -304)
            _ = fi53
        }
        _ = fi48
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v45
    _ = v47
    _ = v5
}
