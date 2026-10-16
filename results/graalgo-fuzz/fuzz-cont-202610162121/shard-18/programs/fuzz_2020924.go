package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int16 = -499
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    _ = len(v2)
    v1 = (v1 << ((0) & 63))
    delete(v3, 479)
    var v4 int32 = 2147483647
    var v5 string = ("a" + "a")
    var v6 int = int(v1)
    var v7 int = (v6 >> ((801) & 63))
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int8 = 100
        if ((!(v5 < (v5 + v5))) && false) {
            var v10 string = "a"
            _ = v10
        }
        fmt.Println(v4)
        _ = fi8
        _ = v9
    }
    v11 := GenS1{F0: "a"}
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v7, v5, v4)
        _ = fi12
    }
    switch ("a" + v5) {
    case "\n\t":
        var v13 string = v5
        _ = v13
    case "foo bar":
        v6 = (v6 + -1)
    }
    var v14 int64 = -764
    for fi15 := 0; fi15 < 9; fi15++ {
        switch v7 {
        case 1:
            v4 -= 42
        default:
            if false {
                v7 -= (fi15 ^ (v6 & int(v1)))
                v6 += (v7 % ((127) | 1))
                var v16 uint32 = 2
                _ = v16
            } else {
                for fi17 := 0; fi17 < 16; fi17++ {
                    switch (fi17 / ((v6) | 1)) {
                    case 32767:
                        var v18 float64 = float64(float32(1.5))
                        _ = v18
                    case 9223372036854775807:
                        v6 -= -450
                    case 42:
                        var v19 uint16 = 0
                        _ = v19
                    }
                    var v20 int8 = -1
                    v21 := GenS2{F0: (668 <= 453), F1: 65535}
                    v22 := GenS2{F0: true, F1: 0}
                    fmt.Println(fi17, v6)
                    v23 := GenS2{F0: (!((3.14159 == 3.14159) || true)), F1: 65535}
                    v24 := []int{fi17, fi17, (v7 | v6), (fi17 - (v7 >> ((684) & 63)))}
                    _ = fi17
                    _ = v20
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                v25 := []int{(v7 - (^(v6)))}
                v7 -= v7
                _ = v25
            }
        }
        v26 := GenS2{F0: false, F1: 0}
        _ = fi15
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
