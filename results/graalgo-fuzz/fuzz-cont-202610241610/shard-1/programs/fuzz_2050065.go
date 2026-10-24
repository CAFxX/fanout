package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint64) int {
    var v6 uint16 = 42
    fmt.Println(v6, p0)
    v6 += (^(v6))
    if (0 <= 0) {
        var v7 int64 = -21
        if (!true) {
            var v8 uint32 = 4294967295
            _ = v8
        } else {
            p0 = p0
        }
        v9 := GenS1{F0: -1, F1: 65535, F2: 32767}
        _ = v7
        _ = v9
    } else {
        var v10 int32 = 2147483647
        v6 = (v6 & v6)
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: 0, F1: 18446744073709551615, F2: (fi11 ^ fi11)}
            fmt.Println(v6)
            if ((-2.25 - 1e10) > 3.14159) {
                v13 := GenS1{F0: 42, F1: 255, F2: (fi11 >> ((0) & 63))}
                v10 = 0
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    _ = p0
    _ = v6
    return 32767
}
func genFn1(p0 string, p1 string, p2 uint16) uint8 {
    var v14 uint8 = 127
    v15 := []int{32, -9223372036854775808, -9223372036854775808}
    v16 := GenS2{F0: false, F1: 2147483647, F2: 353}
    switch (p1 + p0) {
    case "hello":
        var v17 uint64 = 1
        _ = v17
    case "\n\t":
        var v18 int16 = 0
        _ = v18
    case "a":
        if true {
            if ((2147483648 < 2147483647) && (!true)) {
                p1 = "a"
                fmt.Println(p0, p2)
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch fi19 {
                    case 42:
                        fmt.Println(p0, p2)
                    case 127:
                        fmt.Println(p0, p1, p0)
                    }
                    if false {
                        var v20 int64 = -111
                        p0 = p1
                        _ = v20
                    } else {
                        var v21 uint64 = 667
                        _ = v21
                    }
                    if (!false) {
                        p0 = p1
                    }
                    _ = fi19
                }
            }
        }
    default:
        fmt.Println(p1)
    }
    v14 = v14
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return v14
}

func main() {
    var v1 bool = (true || true)
    v2 := make(map[int]int)
    var v3 uint = 42
    v4 := make(map[uint32]int)
    v5 := []string{"\n\t", ("hello" + "foo bar")}
    switch 32767 {
    case 32768:
        for fi22 := 0; fi22 < 7; fi22++ {
            v1 = (!(!false))
            var v23 int = 0
            _ = fi22
            _ = v23
        }
    case 407:
        v1 = false
    case -372:
        v3 = v3
    }
    v1 = (614 == 2)
    v1 = v1
    var v24 float32 = 0.1
    v25 := GenS3{F0: 32767, F1: 2}
    if v1 {
        var v26 uint32 = 0
        _ = v26
    } else {
        var v27 uint8 = genFn1(("hello" + "a"), "hello", 65535)
        if v1 {
            v3 = (v3 << ((1) & 63))
            if v1 {
                v28 := GenS1{F0: 42, F1: v3, F2: -264}
                fmt.Println(v3, v1, v27)
                v24 = float32(0.1)
                _ = v28
            } else {
                if v1 {
                    v24 = 0.0
                    v29 := GenS4{F0: -492, F1: genFn1("a", "\n\t", 65535)}
                    fmt.Println(v1, v24, v24)
                    _ = v29
                }
                v3 = 4294967295
            }
            var v30 int = 32768
            _ = v30
        }
        _ = v27
    }
    v3 -= 255
    var v31 int = -410
    var v32 uint16 = 1
    fmt.Println(v32)
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 16; fi34++ {
            v35 := GenS4{F0: -1, F1: 1}
            var v36 string = (("hello" + "\n\t") + ("foo bar" + "a"))
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    if len(v5) > 0 {
        v5[0] = ""
    }
    var v37 float32 = (float32(3.14159) * 1e10)
    if (0 < (^(v3))) {
        var v38 int16 = 32767
        _ = v38
    } else {
        v39 := GenS4{F0: -741, F1: 252}
        _ = v39
    }
    var v40 bool = (!(!v1))
    v1 = (658 != 42)
    if (v1 || ((!(!true)) && v1)) {
        var v41 int64 = -9223372036854775808
        fmt.Println(v37, v3)
        v41 = -9223372036854775808
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
