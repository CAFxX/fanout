package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = -2.25
    var v2 int16 = 0
    var v3 int32 = 1
    for fi4 := 0; fi4 < 14; fi4++ {
        if (2147483648 == 2147483647) {
            if false {
                v2 += (v2 ^ -32768)
                v5 := GenS1{F0: 9223372036854775807}
                var v6 int = fi4
                _ = v5
                _ = v6
            } else {
                for fi7 := 0; fi7 < 14; fi7++ {
                    switch 249 {
                    case 0:
                        v3 -= v3
                    }
                    v3 += -2147483648
                    var v8 uint8 = 0
                    v9 := []int{fi4, fi4, fi4}
                    var v10 float64 = (v1 + v1)
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            }
            v2 = v2
            switch "hello" {
            case "\n\t":
                v2 -= v2
            }
        } else {
            v1 = (3.14159 * v1)
        }
        v3 += v3
        _ = fi4
    }
    var v11 bool = false
    if (-22 < -9223372036854775808) {
        var v12 string = "hello"
        if (42 < 0) {
            if v11 {
                v12 = v12
            } else {
                v13 := []int{1, 0}
                v13 = append(v13, -511)
                v1 = v1
                _ = v13
            }
        }
        fmt.Println(v2)
        _ = v12
    } else {
        var v14 uint16 = 1
        _ = v14
    }
    v15 := GenS1{F0: 83}
    v2 = (v2 & -1)
    v1 = 0.0
    v3 = v3
    var v16 int8 = 60
    if false {
        for fi17 := 0; fi17 < 1; fi17++ {
            v3 += -1
            v18 := GenS2{F0: ""}
            _ = fi17
            _ = v18
        }
        if v11 {
            for fi19 := 0; fi19 < 11; fi19++ {
                v16 -= v16
                _ = fi19
            }
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                v21 := []int{0}
                _ = fi20
                _ = v21
            }
            v3 = v3
        }
        if (!true) {
            for fi22 := 0; fi22 < 13; fi22++ {
                if (!v11) {
                    var v23 uint16 = 255
                    _ = v23
                }
                if (!v11) {
                    var v24 int16 = v2
                    _ = v24
                }
                v25 := []int{(fi22 / ((32767) | 1)), fi22, (^(fi22))}
                _ = fi22
                _ = v25
            }
            v3 -= (v3 >> ((50) & 63))
            v26 := []int{548}
            _ = v26
        } else {
            for fi27 := 0; fi27 < 9; fi27++ {
                v16 = v16
                _ = fi27
            }
        }
    } else {
        v16 = v16
    }
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
