package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint64 }
func genFn0(p0 int8) uint {
    fmt.Println(p0)
    fmt.Println(p0)
    for fi4 := 0; fi4 < 3; fi4++ {
        fmt.Println(fi4, fi4)
        _ = fi4
    }
    _ = p0
    return 444
}
func genFn1() string {
    v5 := []int{-2, 129}
    if len(v5) > 0 {
        v5[0] = 32767
    }
    v6 := GenS1{F0: 32768, F1: -1}
    if (18446744073709551615 == 1) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 float32 = float32((0.0 + float64(1e10)))
            if false {
                for fi9 := 0; fi9 < 5; fi9++ {
                    fmt.Println(v8, fi9, v8)
                    v8 = 0.0
                    _ = fi9
                }
                switch ("" + "a") {
                case "hello":
                    var v10 string = "hello"
                    _ = v10
                case "":
                    var v11 uint32 = 2147483648
                    _ = v11
                }
                _ = len(v5)
            } else {
                var v12 int8 = 100
                v13 := GenS2{F0: 541, F1: v12}
                _ = v12
                _ = v13
            }
            var v14 int64 = -270
            _ = fi7
            _ = v14
            _ = v8
        }
        v15 := GenS3{F0: 210, F1: 0, F2: 225}
        _ = v15
    }
    var v16 string = "\n\t"
    _ = v16
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v3[4294967295] = 128
    if (62 >= 65535) {
        var v17 float32 = 0.1
        switch 1 {
        case 32768:
            v17 += (1.5 + ((float32(1e10) + v17) * (v17 - v17)))
        case 9223372036854775807:
            v18 := []int{42}
            _ = v18
        }
        _ = v17
    } else {
        var v19 float64 = (float64(0.1) * 0.1)
        if (((0.1 / 1.0) != 1.5) || (-2147483648 != 1)) {
            v19 = v19
        }
        fmt.Println(v19)
        _ = v19
    }
    _ = len(v2)
    if false {
        var v20 bool = (-692 == 530)
        _ = v20
    } else {
        v21 := []int{-824}
        var v22 string = ("hello" + "hello")
        _ = v21
        _ = v22
    }
    var v23 uint8 = 0
    v24 := []int{-1, 769, 208}
    fmt.Println(v23)
    fmt.Println(v23)
    for fi25 := 0; fi25 < 4; fi25++ {
        v23 -= (v23 ^ 227)
        v26 := GenS1{F0: (fi25 >> ((42) & 63)), F1: 66}
        v27 := GenS2{F0: 2, F1: 73}
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint16 = 65535
    if false {
        v28 = v28
        if false {
            var v29 string = ((((("\n\t" + "hello") + ("\n\t" + ("\n\t" + "hello"))) + "hello") + "") + ("hello" + "foo bar"))
            _ = v29
        }
        var v30 uint64 = 615
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
}
