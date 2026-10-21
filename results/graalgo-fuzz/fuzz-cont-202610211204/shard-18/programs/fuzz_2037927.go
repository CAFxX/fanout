package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (float64(1.5) * 0.0)
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    var v5 float32 = 3.14159
    fmt.Println(v5, v5)
    var v6 int16 = 2
    var v7 int64 = -1
    v6 = (v6 / (((v6 | 1)) | 1))
    v5 += 0.0
    v8 := GenS1{F0: 42, F1: 637}
    var v9 bool = true
    for fi10 := 0; fi10 < 6; fi10++ {
        v2 += -2.25
        var v11 uint32 = 2147483648
        if false {
            v11 = v11
            var v12 int8 = 127
            switch ((("a" + ("foo bar" + "foo bar")) + ((((("" + ((("foo bar" + (("foo bar" + "a") + "hello")) + "") + (("hello" + ("" + ("\n\t" + (((("\n\t" + (("a" + "\n\t") + "foo bar")) + "foo bar") + "hello") + "\n\t")))) + ""))) + "hello") + "hello") + "hello") + "a")) + ("a" + ("a" + ("hello" + ((("" + "hello") + "\n\t") + ""))))) {
            case "":
                v11 -= 1
            default:
                for fi13 := 0; fi13 < 3; fi13++ {
                    switch 32767 {
                    case -9223372036854775808:
                        v9 = (!v9)
                    case 0:
                        var v14 int32 = 2147483647
                        _ = v14
                    case 127:
                        v7 -= 1
                    default:
                        fmt.Println(fi10)
                    }
                    var v15 int64 = int64(fi10)
                    _ = fi13
                    _ = v15
                }
            }
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v9 = v9
        v7 += 9223372036854775807
        _ = fi16
    }
    v9 = ((v6 >> ((1) & 63)) > (v6 / ((v6) | 1)))
    for fi17 := 0; fi17 < 1; fi17++ {
        v9 = (v9 || (!(!v9)))
        v18 := []int{fi17, (fi17 & 32768), 0, fi17}
        if (2 > 42) {
            v19 := GenS2{F0: 2, F1: 18446744073709551615, F2: v7}
            var v20 float64 = ((3.14159 + v2) + v2)
            _ = v19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    if ("" <= "\n\t") {
        for fi21 := 0; fi21 < 5; fi21++ {
            if ((v6 / ((v6) | 1)) == v6) {
                var v22 uint8 = 1
                _ = v22
            }
            _ = fi21
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            fmt.Println(v6)
            v24 := GenS2{F0: 664, F1: 673, F2: v7}
            for fi25 := 0; fi25 < 16; fi25++ {
                delete(v4, fi23)
                v26 := GenS1{F0: 0, F1: fi23}
                if v9 {
                    v9 = (128 <= 2)
                    var v27 string = "\n\t"
                    v28 := GenS1{F0: 42, F1: fi23}
                    _ = v27
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        delete(v1, 654)
    }
    fmt.Println(v6, v9)
    if ((v7 << ((2) & 63)) >= v7) {
        v2 = ((1e10 - 1e10) + (1e10 + 1.5))
        delete(v1, 220)
        v29 := GenS2{F0: 0, F1: 1, F2: (^(v7))}
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
