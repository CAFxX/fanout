package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint32 }
func genFn0(p0 int32) bool {
    fmt.Println(p0)
    var v5 string = ""
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{(fi6 >> ((400) & 63)), 127}
        var v8 int64 = 42
        fmt.Println(fi6)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return ((!(-32768 <= -32768)) && (4294967295 >= 1))
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]bool)
    var v3 uint8 = 2
    var v4 int64 = 326
    if (1 < 42) {
        for fi9 := 0; fi9 < 11; fi9++ {
            switch "a" {
            case "\n\t":
                v10 := GenS1{F0: uint32(v3)}
                _ = v10
            case "foo bar":
                var v11 string = "a"
                _ = v11
            default:
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 uint32 = 42
                    var v14 uint16 = 639
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            var v15 int8 = int8(fi9)
            _ = fi9
            _ = v15
        }
        v16 := v2[("a" + (("" + (("foo bar" + "") + "\n\t")) + "hello"))]
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint64 = 2
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        v1 += 3.14159
        for fi19 := 0; fi19 < 4; fi19++ {
            v4 -= v4
            v20 := GenS2{F0: "hello"}
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 uint8 = (v3 >> ((710) & 63))
                switch "a" {
                case "hello":
                    v23 := GenS3{F0: (("foo bar" + "") + ("" + ("foo bar" + (("\n\t" + (("\n\t" + "") + (("" + ("hello" + "")) + ("foo bar" + ("hello" + "\n\t"))))) + "a")))), F1: (v4 * v4)}
                    _ = v23
                case "":
                    var v24 int32 = 2147483647
                    _ = v24
                default:
                    v4 = 693
                }
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v3)
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS4{F0: v3, F1: int16(v3), F2: 538}
        _ = fi25
        _ = v26
    }
    v27 := []int{357, 32768, -331, -1}
    var v28 uint16 = 0
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 int16 = -1
        _ = fi29
        _ = v30
    }
    if len(v27) > 0 {
        v27[0] = -2147483648
    }
    v1 = float32(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
