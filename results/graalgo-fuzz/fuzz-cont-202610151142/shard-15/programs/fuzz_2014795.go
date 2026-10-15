package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    var v1 uint32 = 1
    var v2 uint8 = uint8(v1)
    var v3 int8 = -1
    var v4 int8 = -1
    var v5 string = (("\n\t" + "foo bar") + "hello")
    var v6 uint8 = 123
    v4 -= v4
    fmt.Println(v1, v2)
    fmt.Println(v4, v2, v2)
    var v7 uint = 540
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{(fi8 | fi8)}
        v4 += -73
        fmt.Println(v6)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{251, 32768, (fi10 << (((v7 + (v7 * (v7 / ((uint(v6)) | 1))))) & 63))}
        var v12 int32 = -2147483648
        var v13 int8 = v3
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 string = v5
    var v15 bool = (!("foo bar" == (v5 + "\n\t")))
    if (v5 >= v14) {
        for fi16 := 0; fi16 < 18; fi16++ {
            if false {
                v5 = (v14 + (v14 + (v14 + (v14 + v5))))
                v17 := GenS1{F0: (^(v1)), F1: (^(v6))}
                _ = v17
            } else {
                v5 = v5
                for fi18 := 0; fi18 < 16; fi18++ {
                    switch ("\n\t" + v5) {
                    case "a":
                        v4 -= int8(v7)
                    case "":
                        v19 := []int{int(v2)}
                        _ = v19
                    default:
                        fmt.Println(v7, v7, v7)
                    }
                    fmt.Println(v3, v3)
                    v2 -= (v6 ^ (v6 * (v2 / ((v6) | 1))))
                    v20 := GenS1{F0: 2, F1: (v2 - v2)}
                    var v21 uint64 = 18446744073709551615
                    var v22 uint64 = v21
                    _ = fi18
                    _ = v20
                    _ = v21
                    _ = v22
                }
                var v23 float64 = -2.25
                _ = v23
            }
            v24 := GenS1{F0: (v1 * 0), F1: (v2 & uint8(v7))}
            v2 = 151
            _ = fi16
            _ = v24
        }
        v15 = v15
    } else {
        for fi25 := 0; fi25 < 15; fi25++ {
            switch fi25 {
            case 128:
                for fi26 := 0; fi26 < 2; fi26++ {
                    fmt.Println(v3, v1, v6)
                    v27 := []int{(fi26 << ((255) & 63))}
                    _ = fi26
                    _ = v27
                }
            }
            if true {
                switch 32767 {
                case 42:
                    for fi28 := 0; fi28 < 1; fi28++ {
                        var v29 uint32 = v1
                        if false {
                            v30 := GenS1{F0: (v1 * v1), F1: (v2 * (v2 | v6))}
                            var v31 float64 = 1.5
                            _ = v30
                            _ = v31
                        }
                        _ = fi28
                        _ = v29
                    }
                case -778:
                    v32 := GenS2{F0: (v6 % (((v2 + 255)) | 1)), F1: 222}
                    _ = v32
                case 872:
                    var v33 int8 = (v3 + 1)
                    _ = v33
                }
                var v34 bool = (!(2 >= (^(v1))))
                _ = v34
            }
            fmt.Println(v7)
            _ = fi25
        }
        v7 = 2
        var v35 float32 = 3.14159
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
