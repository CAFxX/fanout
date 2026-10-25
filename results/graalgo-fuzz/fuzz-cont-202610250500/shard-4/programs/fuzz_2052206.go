package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 string; F2 int }
func genFn0(p0 uint) uint {
    var v4 uint32 = 916
    p0 = 4294967295
    _ = p0
    _ = v4
    return (p0 + 1)
}
func genFn1(p0 int8) int32 {
    fmt.Println(p0)
    p0 = -128
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 string = ""
        var v7 uint16 = 40
        v8 := GenS1{F0: ((v6 + v6) + v6), F1: 4294967295}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 int32 = -1
    v3 += v3
    if ((-2.25 / (3.14159 + ((3.14159 - -2.25) / 1.0))) >= 1.5) {
        v3 -= v3
        var v9 string = ("foo bar" + "hello")
        for fi10 := 0; fi10 < 11; fi10++ {
            v3 -= (v3 / (((v3 | (v3 >> ((65535) & 63)))) | 1))
            _ = fi10
        }
        _ = v9
    }
    var v11 int32 = v3
    v11 += 2147483647
    v12 := GenS1{F0: "", F1: 2147483647}
    for fi13 := 0; fi13 < 18; fi13++ {
        switch (fi13 % (((fi13 % (((fi13 / ((32767) | 1))) | 1))) | 1)) {
        case 2147483647:
            v14 := GenS2{F0: -128, F1: 9223372036854775807}
            _ = v14
        case -2147483648:
            switch 4294967295 {
            case 850:
                v3 = (-(v3))
            case 127:
                var v15 string = ("\n\t" + ("hello" + "\n\t"))
                _ = v15
            default:
                var v16 bool = false
                _ = v16
            }
        case -9223372036854775808:
            switch fi13 {
            case 427:
                if (0 > 834) {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        _ = len(v2)
                        _ = fi17
                    }
                    v18 := []int{(fi13 >> ((4294967295) & 63)), (fi13 % (((fi13 * (fi13 - 0))) | 1))}
                    _ = v18
                }
            case 128:
                var v19 int64 = -231
                _ = v19
            }
        default:
            var v20 int32 = v11
            _ = v20
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v3 = (v3 | v11)
            _ = fi21
        }
        var v22 uint8 = 224
        _ = fi13
        _ = v22
    }
    v23 := GenS3{F0: 0}
    switch 260 {
    case 1:
        v24 := []int{638}
        _ = v24
    case 32768:
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := []int{fi25, fi25}
            _ = fi25
            _ = v26
        }
    case 32767:
        v11 += (v11 >> ((65535) & 63))
    default:
        v27 := GenS4{F0: 1000, F1: (((("\n\t" + "\n\t") + "a") + ("" + "foo bar")) + (("foo bar" + "hello") + "")), F2: -532}
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
}
