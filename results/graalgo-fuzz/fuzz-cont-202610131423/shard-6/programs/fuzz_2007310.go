package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
func genFn0(p0 int16, p1 uint8, p2 string) int64 {
    v4 := []int{-778, 32767, 32768}
    var v5 int8 = 127
    fmt.Println(p2)
    for fi6 := 0; fi6 < 18; fi6++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 int64 = 993
    v2 := []string{("hello" + "")}
    v3 := make(map[int]int)
    v1 = v1
    v7 := GenS1{F0: false, F1: -1}
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v8 int64 = (v1 << ((42) & 63))
    v3[-1] = -818
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint16 = 2
        if (42 > int16(fi9)) {
            if false {
                switch 32767 {
                case 0:
                    v8 = v1
                case 128:
                    switch fi9 {
                    case -2147483648:
                        v10 -= 2
                    case 0:
                        v1 = (v8 % (((v8 % (((v8 - v1)) | 1))) | 1))
                    default:
                        var v11 int = (fi9 / (((fi9 * fi9)) | 1))
                        _ = v11
                    }
                }
            } else {
                var v12 int32 = 42
                v8 += (v1 - (v1 >> ((42) & 63)))
                var v13 float64 = 0.0
                _ = v12
                _ = v13
            }
            v1 -= (v1 % ((genFn0(1, 160, ("hello" + "foo bar"))) | 1))
            delete(v3, fi9)
        } else {
            var v14 bool = (-128 <= -1)
            var v15 int8 = 1
            if v14 {
                var v16 uint8 = 255
                _ = v16
            } else {
                switch (("hello" + "hello") + ((("" + "\n\t") + (("" + "") + ("hello" + "\n\t"))) + "hello")) {
                case "hello":
                    if v14 {
                        for fi17 := 0; fi17 < 15; fi17++ {
                            var v18 float64 = (1.5 / 1.0)
                            _ = fi17
                            _ = v18
                        }
                    }
                default:
                    fmt.Println(v1, v10)
                }
                v1 += -908
                v10 += v10
            }
            _ = v14
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    v19 := GenS1{F0: false, F1: -1}
    v20 := GenS1{F0: (2 == 610), F1: 32767}
    v1 = genFn0(0, 0, (("hello" + "\n\t") + "foo bar"))
    switch "foo bar" {
    case "":
        if (255 >= 162) {
            switch -2147483648 {
            case 116:
                v21 := GenS1{F0: ((65535 == 18446744073709551615) || (65535 > 2)), F1: -1}
                _ = v21
            case -849:
                for fi22 := 0; fi22 < 2; fi22++ {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        if len(v2) > 0 {
                            v2[0] = ("a" + "foo bar")
                        }
                        _ = fi23
                    }
                    _ = len(v2)
                    _ = fi22
                }
            }
            v24 := GenS1{F0: ((-9223372036854775808 > 42) || (-1 > -345)), F1: 882}
            for fi25 := 0; fi25 < 1; fi25++ {
                fmt.Println(v1, v1, v1)
                _ = fi25
            }
            _ = v24
        }
    case "foo bar":
        if true {
            var v26 string = ("\n\t" + "")
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
}
