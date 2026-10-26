package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 bool }
func genFn0() int8 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint16 = 162
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 255, F1: -1, F2: 42}
        if (0 > 128) {
            switch (((("\n\t" + (("a" + (("hello" + "foo bar") + ("a" + ("hello" + ("" + "\n\t"))))) + ("a" + ("a" + ("a" + "\n\t"))))) + "a") + "\n\t") + "a") {
            case "hello":
                for fi10 := 0; fi10 < 17; fi10++ {
                    if true {
                        var v11 uint32 = 2147483648
                        var v12 int8 = 42
                        _ = v11
                        _ = v12
                    }
                    v13 := GenS2{F0: 255}
                    var v14 uint = 1
                    var v15 uint32 = 4294967295
                    _ = fi10
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
        }
        _ = fi8
        _ = v9
    }
    v16 := []int{}
    _ = v16
    return 94
}

func main() {
    var v1 int8 = 127
    var v2 uint = 0
    v3 := []int{-1, 0}
    v4 := []string{"hello", "hello", "foo bar"}
    v5 := make(map[int]string)
    fmt.Println(v1, v1)
    var v17 int = -2147483648
    fmt.Println(v17)
    if (0 != 875) {
        for fi18 := 0; fi18 < 20; fi18++ {
            for fi19 := 0; fi19 < 15; fi19++ {
                v17 = (fi18 / ((fi18) | 1))
                delete(v5, fi18)
                v1 -= -41
                _ = fi19
            }
            v5[1] = ""
            _ = fi18
        }
        v20 := GenS3{F0: false, F1: 169, F2: (v2 >= v2)}
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 string = "\n\t"
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        fmt.Println(v1)
    }
    var v23 int64 = -71
    var v24 int8 = (v1 + (v1 ^ (-(v1))))
    var v25 int = (v17 << ((255) & 63))
    var v26 int8 = v24
    switch (v25 % (((v25 / ((-744) | 1))) | 1)) {
    case -1:
        switch (^(v25)) {
        case 4294967295:
            var v27 int = v25
            _ = v27
        case 936:
            delete(v5, v17)
        case 32768:
            fmt.Println(v2, v24, v26)
        }
    case -457:
        fmt.Println(v1)
    case 128:
        v28 := GenS1{F0: 255, F1: 42, F2: v2}
        _ = v28
    default:
        if false {
            var v29 float32 = -2.25
            fmt.Println(v25, v26)
            var v30 uint8 = 133
            _ = v29
            _ = v30
        } else {
            v31 := v5[v17]
            v17 -= (v17 + v17)
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
