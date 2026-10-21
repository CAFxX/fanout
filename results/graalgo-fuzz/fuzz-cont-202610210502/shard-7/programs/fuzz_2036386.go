package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 string; F1 int32; F2 int8 }

func main() {
    var v1 string = (("\n\t" + ("hello" + "\n\t")) + "a")
    var v2 uint16 = 151
    var v3 int32 = 1
    if (true && false) {
        v4 := GenS1{F0: 32767, F1: 1}
        if (!(((65535 > 969) && true) || (!false))) {
            for fi5 := 0; fi5 < 12; fi5++ {
                if false {
                    v3 -= v3
                } else {
                    for fi6 := 0; fi6 < 2; fi6++ {
                        v2 += (v2 ^ 65535)
                        v7 := []int{fi5}
                        fmt.Println(v2, v1)
                        _ = fi6
                        _ = v7
                    }
                    v3 = v3
                    for fi8 := 0; fi8 < 16; fi8++ {
                        v9 := []int{(-(fi5))}
                        v10 := GenS2{F0: 2, F1: (255 != 2), F2: 0}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := GenS2{F0: 18446744073709551615, F1: (float64(1e10) <= -2.25), F2: 1}
                    _ = fi11
                    _ = v12
                }
                _ = fi5
            }
            v13 := []int{}
            _ = v13
        }
        _ = v4
    }
    v14 := GenS3{F0: (756 == 42), F1: 42}
    var v15 uint = 255
    v3 -= 42
    v16 := GenS4{F0: v1, F1: v3, F2: 100}
    var v17 uint = v15
    var v18 int8 = 127
    var v19 int32 = 0
    v15 = 401
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
