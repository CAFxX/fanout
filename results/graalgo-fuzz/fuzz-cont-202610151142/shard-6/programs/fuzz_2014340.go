package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int8 }
func genFn0(p0 bool, p1 uint8, p2 int32) int8 {
    p0 = p0
    var v7 bool = p0
    fmt.Println(v7)
    v8 := []int{307, 210}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}
func genFn1() bool {
    v9 := GenS1{F0: 241, F1: "hello", F2: -1}
    var v10 float32 = 0.1
    _ = v10
    _ = v9
    return (-1 != 2147483647)
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{311}
    v3 := make(map[int]int)
    var v4 uint16 = 65535
    var v5 int32 = -911
    v6 := []uint32{2147483648}
    var v11 uint64 = 18446744073709551615
    if false {
        fmt.Println(v4, v4)
        if (127 >= 174) {
            v2 = append(v2, -2147483648)
            var v12 int16 = 32767
            _ = v12
        }
    } else {
        v11 = (v11 << ((4294967295) & 63))
        var v13 uint = 2
        var v14 float64 = -2.25
        _ = v13
        _ = v14
    }
    var v15 uint16 = v4
    v15 -= v4
    for fi16 := 0; fi16 < 8; fi16++ {
        v11 = 42
        delete(v1, "")
        _ = fi16
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        v17 := v1["hello"]
        _ = v17
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = 0
        }
    case "a":
        switch 432 {
        case -237:
            if ((true || true) || false) {
                var v18 int64 = -307
                _ = v18
            } else {
                var v19 uint16 = uint16(v11)
                var v20 bool = false
                v1[("foo bar" + "a")] = 1
                _ = v19
                _ = v20
            }
        case 1:
            var v21 float32 = 3.14159
            _ = v21
        case 0:
            v4 = 556
        default:
            var v22 bool = false
            _ = v22
        }
    default:
        var v23 int16 = -258
        _ = v23
    }
    switch 380 {
    case 32767:
        delete(v1, "hello")
    case 9223372036854775807:
        var v24 float32 = -2.25
        _ = v24
    default:
        var v25 float64 = 3.14159
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 int = (fi26 * -148)
        for fi28 := 0; fi28 < 13; fi28++ {
            if false {
                v27 = fi28
                if true {
                    var v29 int32 = (v5 * (v5 >> ((1) & 63)))
                    _ = v29
                } else {
                    if len(v6) > 0 {
                        v6[0] = 2147483648
                    }
                    var v30 float32 = 0.1
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v30
                }
                v31 := GenS1{F0: 180, F1: "\n\t", F2: 63}
                _ = v31
            } else {
                switch ((("hello" + "a") + "hello") + "\n\t") {
                case "\n\t":
                    switch "a" {
                    case "":
                        var v32 float64 = 0.1
                        _ = v32
                    case "a":
                        v15 = v4
                        v27 = fi28
                        fmt.Println(fi28, fi28, v4)
                    case "foo bar":
                        v33 := GenS1{F0: uint8(v15), F1: "hello", F2: 42}
                        _ = v33
                    default:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                case "a":
                    v34 := GenS1{F0: 128, F1: "hello", F2: 0}
                    _ = v34
                default:
                    if (((380 <= 2) || (!(!(("hello" >= ("" + ("foo bar" + "a"))) || (true && false))))) || (42 == -9223372036854775808)) {
                        fmt.Println(v5)
                    } else {
                        v27 = (fi26 << ((653) & 63))
                        v35 := GenS2{F0: -49}
                        fmt.Println(fi26, fi26, v15)
                        _ = v35
                    }
                }
            }
            _ = fi28
        }
        switch "" {
        case "foo bar":
            fmt.Println(v15)
        case "hello":
            if false {
                fmt.Println(v15, v4, v4)
            }
        case "":
            var v36 bool = false
            _ = v36
        default:
            switch v27 {
            case 9223372036854775807:
                _ = len(v2)
            default:
                v3[v27] = fi26
            }
        }
        _ = fi26
        _ = v27
    }
    v5 = (v5 << ((65535) & 63))
    v11 = 700
    v11 = v11
    v11 = 726
    v37 := GenS3{F0: 440}
    var v38 uint32 = 65535
    v5 = 0
    v3[9223372036854775807] = 483
    v39 := GenS4{F0: v11, F1: (v11 ^ (^(v11))), F2: 106}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
