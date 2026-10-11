package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint32 }
func genFn0() bool {
    var v5 int16 = 32767
    var v6 uint32 = 2147483648
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 bool, p1 uint8, p2 int) string {
    switch 2147483647 {
    case -32:
        v7 := GenS1{F0: p1}
        _ = v7
    }
    fmt.Println(p1, p0)
    if false {
        v8 := GenS1{F0: p1}
        var v9 int32 = 0
        v10 := GenS1{F0: p1}
        _ = v10
        _ = v8
        _ = v9
    }
    switch 4294967295 {
    case -333:
        v11 := []int{(p2 + p2), p2, (p2 | p2)}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return ("hello" + "a")
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 1e10
    var v3 float32 = ((1e10 + 1.5) + 1.5)
    var v4 int8 = 100
    v2 = v2
    switch "" {
    case "\n\t":
        if true {
            var v12 float32 = v3
            _ = v12
        }
    case "a":
        v4 = v4
    case "":
        v2 += v2
    default:
        v13 := GenS1{F0: 1}
        _ = v13
    }
    var v14 float64 = (-2.25 * v2)
    v15 := []int{901, -1, -989, 250}
    var v16 string = (("" + ((("" + "hello") + "\n\t") + "\n\t")) + "a")
    if false {
        var v17 bool = true
        _ = v17
    } else {
        v3 = v3
        for fi18 := 0; fi18 < 17; fi18++ {
            if false {
                v19 := GenS2{F0: v4}
                _ = v19
            }
            for fi20 := 0; fi20 < 3; fi20++ {
                switch (fi18 - (^(fi18))) {
                case -1:
                    v21 := GenS1{F0: 42}
                    _ = v21
                case -624:
                    v22 := GenS3{F0: 32767, F1: 2147483647}
                    v14 -= -2.25
                    _ = v22
                }
                _ = fi20
            }
            _ = fi18
        }
    }
    var v23 int = 0
    v3 -= -2.25
    if false {
        v24 := GenS2{F0: (v4 / ((v4) | 1))}
        fmt.Println(v14)
        _ = v24
    }
    var v25 uint8 = 70
    var v26 uint32 = 2147483648
    for fi27 := 0; fi27 < 2; fi27++ {
        fmt.Println(v2, v16, v14)
        var v28 bool = (((-(v4)) < (v4 - 1)) || (true && (((!true) || (true || (true || (153 >= 1)))) && (!(!(true && true))))))
        var v29 float32 = v3
        _ = fi27
        _ = v28
        _ = v29
    }
    if len(v15) > 0 {
        v15[0] = int(v4)
    }
    v30 := GenS2{F0: -1}
    v23 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
}
