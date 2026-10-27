package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }

func main() {
    v1 := []int{-9223372036854775808}
    var v2 float32 = (1.5 / 0.1)
    var v3 float64 = 0.1
    switch 128 {
    case 4294967295:
        _ = len(v1)
    case 0:
        for fi4 := 0; fi4 < 7; fi4++ {
            if true {
                if (-125 < 1) {
                    v3 = -2.25
                    var v5 int = 9223372036854775807
                    _ = len(v1)
                    _ = v5
                }
                v1 = append(v1, fi4)
                if false {
                    v6 := GenS1{F0: false, F1: -1}
                    for fi7 := 0; fi7 < 17; fi7++ {
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            var v8 float32 = ((-2.25 / v2) / (v2 + 0.0))
                            var v9 string = ""
                            v3 = 3.14159
                            _ = v8
                            _ = v9
                        }
                        _ = fi7
                    }
                    if (((94 != 95) && (-804 == 0)) || false) {
                        if true {
                            v10 := GenS1{F0: (v3 == ((1.5 - -2.25) / (v3 * (v3 - v3)))), F1: 2147483647}
                            _ = v10
                        }
                        var v11 int16 = 0
                        _ = v11
                    }
                    _ = v6
                }
            }
            _ = fi4
        }
    case 32768:
        var v12 bool = (127 != 1)
        _ = v12
    }
    fmt.Println(v2)
    for fi13 := 0; fi13 < 13; fi13++ {
        v2 = v2
        _ = fi13
    }
    fmt.Println(v2)
    for fi14 := 0; fi14 < 9; fi14++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi14
    }
    var v15 uint = 42
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS1{F0: true, F1: 0}
        _ = fi16
        _ = v17
    }
    var v18 uint = (v15 ^ v15)
    v2 = v2
    var v19 float64 = ((v3 * ((3.14159 + (float64(v2) - (v3 + 3.14159))) * -2.25)) + v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
