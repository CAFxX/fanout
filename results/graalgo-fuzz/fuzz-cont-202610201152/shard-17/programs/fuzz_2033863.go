package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 32767
    var v3 int8 = 0
    var v4 float64 = (1.5 + 1.5)
    var v5 int = -668
    v6 := make(map[uint32]bool)
    var v7 uint16 = (v2 % ((v2) | 1))
    var v8 int8 = -1
    for fi9 := 0; fi9 < 2; fi9++ {
        v3 -= 0
        v7 = v2
        _ = fi9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v7 = v7
        _ = fi10
    }
    v11 := GenS1{F0: (!((2147483647 <= -2147483648) && false))}
    delete(v1, "")
    var v12 uint16 = (v7 / ((0) | 1))
    var v13 float32 = float32((((((v4 + v4) * ((v4 * (v4 - ((1.5 + float64(3.14159)) - v4))) + 1.5)) / -2.25) + v4) / (v4 - float64((0.0 / float32(3.14159))))))
    fmt.Println(v5)
    var v14 int = v5
    v15 := GenS1{F0: (!((!(7 >= 124)) || (!(-505 > 1))))}
    fmt.Println(v8)
    var v16 string = ((("\n\t" + "a") + "hello") + "")
    if ((v13 * -2.25) == 1e10) {
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 int = 32767
            if true {
                var v19 string = v16
                _ = v19
            } else {
                v20 := []int{v5, 2147483647}
                v20 = append(v20, v18)
                if false {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v8 = 19
                        v12 = v2
                        _ = fi21
                    }
                    v4 -= 1e10
                }
                _ = v20
            }
            if true {
                v7 += v12
                for fi22 := 0; fi22 < 18; fi22++ {
                    v4 += v4
                    v23 := []int{42, (v14 >> ((42) & 63))}
                    v24 := GenS1{F0: false}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v4 = v4
            } else {
                var v25 uint32 = 42
                var v26 uint16 = (v2 | (^(v2)))
                _ = v25
                _ = v26
            }
            _ = fi17
            _ = v18
        }
    }
    v27 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
