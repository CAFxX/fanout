package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint64, p1 bool) uint8 {
    for fi4 := 0; fi4 < 1; fi4++ {
        fmt.Println(p0, p0)
        for fi5 := 0; fi5 < 9; fi5++ {
            switch fi4 {
            case 32767:
                v6 := GenS1{F0: 42}
                _ = v6
            }
            var v7 int64 = 0
            _ = fi5
            _ = v7
        }
        var v8 float32 = (1e10 * (((3.14159 + 1.5) * 1e10) + ((0.1 - 0.0) / 1.0)))
        _ = fi4
        _ = v8
    }
    p1 = p1
    v9 := []int{32768, -1, 667, -490}
    _ = p0
    _ = p1
    _ = v9
    return 98
}
func genFn1() int64 {
    var v10 string = "a"
    var v11 uint32 = 831
    _ = v10
    _ = v11
    return -9223372036854775808
}

func main() {
    v1 := []int{}
    var v2 bool = (true || ((!(94 > -344)) || ((0 == 2147483647) && (989 > 9223372036854775807))))
    var v3 float64 = 0.0
    if (!(v2 && (1.5 == (1e10 + 0.0)))) {
        v1 = append(v1, 1)
        if (!(574 == 155)) {
            var v12 uint64 = 1
            _ = v12
        } else {
            if v2 {
                if v2 {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v3 = v3
                        _ = fi13
                    }
                    _ = len(v1)
                }
            } else {
                _ = len(v1)
                var v14 uint8 = 1
                if len(v1) > 0 {
                    v1[0] = -2147483648
                }
                _ = v14
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v15 int16 = 32767
        _ = v15
    } else {
        v2 = (2 < 2)
        var v16 string = ""
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint16 = 1
        _ = fi17
        _ = v18
    }
    var v19 uint16 = 382
    _ = len(v1)
    v20 := GenS1{F0: -9223372036854775808}
    v3 = ((float64(float32((v3 / 1.0))) / v3) + v3)
    var v21 uint = 18446744073709551615
    var v22 int32 = int32(v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
