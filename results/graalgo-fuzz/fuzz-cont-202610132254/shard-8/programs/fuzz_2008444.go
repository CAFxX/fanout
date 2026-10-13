package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string; F1 string }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (float64((1.5 - (3.14159 - (0.1 - 1.5)))) < 1e10)
    var v3 uint16 = 36
    if true {
        _ = len(v1)
        var v4 uint32 = 42
        _ = v4
    }
    var v5 int8 = 100
    v6 := v1[-506]
    v7 := GenS1{F0: ("" + ""), F1: 54}
    if true {
        v3 = (^(v3))
        var v8 int16 = 42
        switch -9223372036854775808 {
        case 1:
            fmt.Println(v6, v5)
        default:
            var v9 bool = v2
            _ = v9
        }
        _ = v8
    } else {
        for fi10 := 0; fi10 < 6; fi10++ {
            switch (fi10 ^ int(v5)) {
            case 128:
                v11 := GenS2{F0: ("a" + "hello"), F1: "hello"}
                _ = v11
            case 706:
                var v12 int = fi10
                _ = v12
            }
            for fi13 := 0; fi13 < 10; fi13++ {
                if true {
                    v14 := GenS2{F0: "", F1: "\n\t"}
                    v2 = (!true)
                    _ = v14
                }
                _ = fi13
            }
            _ = fi10
        }
        v1[4294967295] = (!v2)
    }
    var v15 int32 = 42
    if (((float32(0.0) + float32(0.1)) * (0.0 * 1e10)) > -2.25) {
        var v16 bool = v6
        for fi17 := 0; fi17 < 18; fi17++ {
            v5 += (v5 | (v5 / ((127) | 1)))
            if (!(false && v6)) {
                v3 = 255
                v18 := v1[fi17]
                _ = v18
            } else {
                for fi19 := 0; fi19 < 1; fi19++ {
                    delete(v1, fi19)
                    v5 = v5
                    _ = fi19
                }
                v15 = (^(v15))
            }
            var v20 int16 = 1
            _ = fi17
            _ = v20
        }
        _ = v16
    } else {
        var v21 float32 = (((float32(float64(3.14159)) - (-2.25 * float32(0.0))) * ((-2.25 / 1.5) + -2.25)) - -2.25)
        for fi22 := 0; fi22 < 14; fi22++ {
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 bool = ((v2 || (0 < 1)) && false)
                _ = fi23
                _ = v24
            }
            v25 := []int{4294967295, -9223372036854775808}
            _ = fi22
            _ = v25
        }
        _ = v21
    }
    v2 = v2
    v6 = v2
    v15 = (^(v15))
    fmt.Println(v2, v6, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
