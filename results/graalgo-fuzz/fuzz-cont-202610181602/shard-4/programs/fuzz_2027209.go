package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }
func genFn0() int8 {
    var v6 uint8 = 128
    v7 := []int{9223372036854775807, 4294967295, -2147483648}
    _ = v6
    _ = v7
    return -128
}
func genFn1(p0 uint16, p1 uint32) int16 {
    var v8 int = 127
    v9 := []int{v8, -865}
    var v10 uint = 1
    v11 := GenS1{F0: "hello", F1: -1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    var v3 float32 = 0.0
    v4 := make(map[string]string)
    var v5 int8 = 115
    for fi12 := 0; fi12 < 8; fi12++ {
        switch ("hello" + (("" + "foo bar") + "a")) {
        case "foo bar":
            var v13 uint16 = 591
            _ = v13
        case "\n\t":
            var v14 uint32 = 1
            _ = v14
        default:
            v15 := GenS1{F0: (("\n\t" + ("hello" + "a")) + ("hello" + ("hello" + ""))), F1: (v5 & v5)}
            _ = v15
        }
        var v16 bool = true
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 int = 128
            fmt.Println(v5, v3)
            fmt.Println(v18, fi12, fi17)
            _ = fi17
            _ = v18
        }
        _ = fi12
        _ = v16
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := GenS2{F0: 4294967295}
        var v21 int16 = -32768
        if (true && (42 > 162)) {
            v22 := GenS1{F0: (("foo bar" + "hello") + "hello"), F1: (v5 | v5)}
            if ((fi19 < int(v5)) || (!(true || false))) {
                v21 = -736
                v23 := []int{(fi19 - 9223372036854775807), fi19}
                _ = v23
            }
            var v24 string = "\n\t"
            _ = v22
            _ = v24
        } else {
            fmt.Println(v3, v3, fi19)
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v25 := GenS1{F0: "hello", F1: (v5 << ((2) & 63))}
    v26 := GenS1{F0: "a", F1: (v5 * genFn0())}
    v5 += (v5 >> ((18446744073709551615) & 63))
    if (127 <= 244) {
        v27 := []int{9223372036854775807, 2147483647, -2147483648, 290}
        _ = v27
    }
    fmt.Println(v3)
    fmt.Println(v5, v3)
    var v28 uint32 = 921
    switch 982 {
    case 1:
        v28 = v28
    case 42:
        v28 += 65535
    case -1:
        v29 := GenS1{F0: "hello", F1: (v5 * v5)}
        _ = v29
    default:
        if (65535 == v28) {
            var v30 uint8 = uint8(v28)
            v5 = 1
            _ = v30
        } else {
            v2[65535] = 0
            fmt.Println(v3, v5, v3)
            v31 := GenS3{F0: 2, F1: 42, F2: 483}
            _ = v31
        }
    }
    v28 -= (v28 | (v28 + v28))
    v32 := []int{2147483647, 1, 32768}
    v3 = ((3.14159 + (v3 - v3)) + 1.5)
    var v33 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
