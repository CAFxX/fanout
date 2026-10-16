package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int16; F1 int32; F2 int8 }
func genFn0(p0 bool, p1 uint32, p2 int16) int32 {
    switch 9223372036854775807 {
    case 32768:
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(p0, p2)
            v7 := []int{(fi6 / ((32767) | 1)), fi6, fi6}
            _ = fi6
            _ = v7
        }
    default:
        v8 := []int{161, -376, 32767, 0}
        _ = v8
    }
    var v9 int8 = int8(p1)
    if (p0 && (65535 == 18446744073709551615)) {
        v10 := GenS1{F0: 982}
        v11 := []int{306, -1, 32767}
        p2 = int16(p1)
        _ = v10
        _ = v11
    }
    var v12 uint16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{2147483648, 180}
    var v2 uint = 2
    v3 := []string{"", (("foo bar" + ("" + "")) + "a"), ("hello" + (("a" + ((("foo bar" + ("hello" + ("a" + "a"))) + "\n\t") + "hello")) + "foo bar"))}
    var v4 int32 = -2147483648
    var v5 uint32 = 2147483647
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint32 = (v5 % ((v5) | 1))
            switch fi13 {
            case 1:
                var v16 int32 = v4
                _ = v16
            case 0:
                var v17 int8 = 126
                _ = v17
            default:
                if true {
                    v18 := GenS2{F0: 42, F1: 65535}
                    _ = v18
                }
            }
            var v19 uint8 = 255
            var v20 uint8 = 1
            _ = fi14
            _ = v15
            _ = v19
            _ = v20
        }
        v2 = v2
        v4 -= (v4 + (v4 ^ int32(fi13)))
        _ = fi13
    }
    fmt.Println(v4)
    var v21 bool = true
    v2 = 779
    v22 := GenS1{F0: genFn0(v21, (^(v5)), 1)}
    var v23 int8 = 42
    v21 = v21
    var v24 int64 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v25 bool = v21
    var v26 uint = v2
    v27 := GenS3{F0: 603, F1: genFn0((("a" + ((((("\n\t" + "hello") + (("hello" + "\n\t") + ("\n\t" + "foo bar"))) + "") + "") + "hello")) != ""), (v5 / ((4294967295) | 1)), int16(v24)), F2: (v23 ^ (v23 / ((v23) | 1)))}
    v28 := GenS3{F0: 0, F1: v4, F2: (v23 << ((v2) & 63))}
    fmt.Println(v24, v4, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
