package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }

func main() {
    var v1 string = "hello"
    var v2 int8 = 100
    var v3 string = "foo bar"
    var v4 uint8 = 0
    v3 = "foo bar"
    var v5 string = ("\n\t" + (v1 + v1))
    var v6 int64 = -933
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint64 = 2
        if false {
            var v9 uint32 = 1
            switch (v3 + v5) {
            case "a":
                for fi10 := 0; fi10 < 1; fi10++ {
                    fmt.Println(v1)
                    v11 := []int{128, 32768, -327, fi7}
                    fmt.Println(v6, v4)
                    var v12 float32 = -2.25
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "":
                for fi13 := 0; fi13 < 12; fi13++ {
                    if (!(false && false)) {
                        if (!false) {
                            v6 -= v6
                            v9 = (v9 % (((v9 + v9)) | 1))
                        }
                        var v14 uint = 1
                        if true {
                            v15 := []int{fi7, (^(fi7)), (-(fi13)), (fi7 << ((296) & 63))}
                            v9 -= (v9 << (((v14 | v14)) & 63))
                            var v16 int8 = -128
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    }
                    v17 := []int{(fi13 >> ((255) & 63)), (fi7 * fi13)}
                    _ = fi13
                    _ = v17
                }
            case "\n\t":
                v18 := GenS1{F0: (v6 * v6), F1: (v9 + v9), F2: int(v2)}
                _ = v18
            }
            v4 = (v4 | 255)
            _ = v9
        }
        v4 -= v4
        _ = fi7
        _ = v8
    }
    var v19 int8 = 0
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 int16 = 0
        _ = fi20
        _ = v21
    }
    v19 = (v19 ^ (v2 - 127))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
