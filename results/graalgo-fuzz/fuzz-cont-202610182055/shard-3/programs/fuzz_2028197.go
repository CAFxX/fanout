package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 string }
type GenS2 struct { F0 int64; F1 int64; F2 string }
type GenS3 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 int32 = 10
    var v4 bool = (1 != 164)
    v5 := make(map[string]int)
    v6 := make(map[uint32]string)
    v3 -= (v3 & v3)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 95, F1: -9, F2: "\n\t"}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 4, F1: -19, F2: ((("\n\t" + "\n\t") + (("a" + "foo bar") + "hello")) + "")}
    fmt.Println(v3, v3, v4)
    switch 4294967295 {
    case 127:
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int8 = -1
            v12 := []int{fi10}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v3)
    if (v4 || true) {
        var v13 bool = (v3 >= v3)
        for fi14 := 0; fi14 < 5; fi14++ {
            v4 = (v2 >= v2)
            if v13 {
                v15 := GenS2{F0: 42, F1: 709, F2: "hello"}
                if (-600 <= -496) {
                    v1[(fi14 - -205)] = 4294967295
                    var v16 uint = 1
                    v3 = -2147483648
                    _ = v16
                }
                v2 -= v2
                _ = v15
            } else {
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 string = ("\n\t" + "a")
                    v19 := GenS3{F0: 1, F1: v2}
                    var v20 uint32 = 2
                    if false {
                        var v21 int64 = 0
                        v18 = ("" + "a")
                        v18 = ("\n\t" + "foo bar")
                        v20 -= 2
                        fmt.Println(v3, fi17, v18)
                        _ = v21
                    }
                    v4 = v4
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            }
            fmt.Println(v2)
            _ = fi14
        }
        fmt.Println(v2, v2, v4)
        _ = v13
    }
    v4 = v4
    if v4 {
        var v22 int32 = v3
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint8 = 2
        _ = fi23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
