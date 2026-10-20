package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int; F2 uint64 }
func genFn0(p0 int64, p1 string, p2 int16) bool {
    v7 := GenS1{F0: 4294967295, F1: -1}
    fmt.Println(p2)
    var v8 uint64 = 42
    var v9 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (!false)
}
func genFn1(p0 uint) int {
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := []int{fi10, fi10, 879, 610}
        v11 = append(v11, fi10)
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 1, F1: 42}
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 string = ""
        v15 := GenS1{F0: 18446744073709551615, F1: 2147483647}
        _ = fi13
        _ = v14
        _ = v15
    }
    if (!(true || true)) {
        v16 := GenS1{F0: 1, F1: 2147483647}
        _ = v16
    }
    _ = p0
    _ = v12
    return 477
}

func main() {
    var v1 bool = (!(0 >= 1))
    var v2 int32 = 0
    v3 := make(map[int]int)
    var v4 float32 = 1.5
    var v5 float64 = 1e10
    var v6 uint8 = 2
    v17 := GenS1{F0: 232, F1: v2}
    var v18 int16 = -1
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS1{F0: 2, F1: v2}
            var v22 string = "hello"
            if (true && false) {
                var v23 int = 2147483647
                _ = v23
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    fmt.Println(v6, v4, v4)
    for fi24 := 0; fi24 < 8; fi24++ {
        v18 = 759
        switch (-(fi24)) {
        case 2147483647:
            switch fi24 {
            case -9223372036854775808:
                var v25 uint64 = 18446744073709551615
                _ = v25
            default:
                if v1 {
                    v18 = -70
                    if (!v1) {
                        var v26 float64 = 1e10
                        _ = v26
                    }
                } else {
                    var v27 uint = 1
                    if v1 {
                        for fi28 := 0; fi28 < 5; fi28++ {
                            v29 := GenS2{F0: 80}
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        v2 -= 1
                    }
                    var v30 float64 = (v5 - v5)
                    _ = v27
                    _ = v30
                }
            }
        case 1:
            var v31 string = ("a" + "\n\t")
            _ = v31
        }
        var v32 int16 = (v18 - (v18 >> ((758) & 63)))
        _ = fi24
        _ = v32
    }
    var v33 uint8 = 0
    fmt.Println(v4, v2, v1)
    v34 := GenS3{F0: (-(v2)), F1: -275, F2: 42}
    v4 = ((1e10 / (-2.25 * 1e10)) - (float32(v5) * (-2.25 + 0.1)))
    v6 += v33
    var v35 bool = true
    var v36 uint8 = v33
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
