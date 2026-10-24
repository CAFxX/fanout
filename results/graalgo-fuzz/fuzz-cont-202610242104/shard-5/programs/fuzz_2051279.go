package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 string }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
func genFn0(p0 int64, p1 int64, p2 int32) int32 {
    if false {
        v5 := []int{-2147483648, -1}
        _ = v5
    }
    v6 := []int{127}
    p1 = (p0 | p1)
    p2 -= -43
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1() uint64 {
    var v7 int = 439
    v7 -= 1
    var v8 uint = 18446744073709551615
    var v9 string = "\n\t"
    v10 := []int{v7, int(v8)}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 2
    v3 := []string{("a" + "")}
    var v4 float32 = -2.25
    if len(v3) > 0 {
        v3[0] = (("a" + "\n\t") + "foo bar")
    }
    fmt.Println(v2, v2)
    v2 = v2
    v11 := v1[665]
    v4 += v4
    v4 -= (float32(0.1) + 0.0)
    if true {
        switch ("a" + ((("foo bar" + "hello") + "\n\t") + ("hello" + "\n\t"))) {
        case "a":
            for fi12 := 0; fi12 < 3; fi12++ {
                v4 = v4
                switch "\n\t" {
                case "":
                    var v13 bool = ((true && (1 == uint8(v11))) || (v2 == 0))
                    _ = v13
                case "foo bar":
                    switch ("a" + ((("hello" + (("" + "\n\t") + "\n\t")) + ("\n\t" + "\n\t")) + (((("\n\t" + "\n\t") + "hello") + "\n\t") + ("" + "a")))) {
                    case "foo bar":
                        v14 := GenS1{F0: (v2 << ((4294967295) & 63)), F1: 2, F2: ("foo bar" + "\n\t")}
                        _ = v14
                    default:
                        if false {
                            v11 = 2147483648
                        }
                    }
                }
                _ = fi12
            }
        case "hello":
            v15 := GenS2{F0: 42, F1: -325, F2: 9223372036854775807}
            _ = v15
        default:
            _ = len(v3)
        }
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            switch fi16 {
            case 32767:
                v4 += v4
            case -186:
                if len(v3) > 0 {
                    v3[0] = "a"
                }
            case -1:
                var v18 int8 = int8(fi17)
                _ = v18
            }
            v19 := v1[(-(fi17))]
            var v20 int = (^(fi17))
            _ = fi17
            _ = v19
            _ = v20
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (1.5 < 1e10) {
            fmt.Println(v11, fi16, v4)
        } else {
            v21 := GenS2{F0: 182, F1: 42, F2: -440}
            fmt.Println(v4, v11, v11)
            var v22 int = fi16
            _ = v21
            _ = v22
        }
        _ = fi16
    }
    v11 = uint32(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
