package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint64, p1 bool) int64 {
    for fi7 := 0; fi7 < 3; fi7++ {
        p0 = p0
        _ = fi7
    }
    v8 := GenS1{F0: (((!p1) || ((42 >= 127) && ((4294967295 >= 1) || p1))) && (!(((!((p1 || ((float32(1e10) >= (3.14159 / (3.14159 / -2.25))) || (-2147483648 > 1))) && (1 < 880))) || ((2147483647 >= 2147483647) || ("a" < ("" + "a")))) && (!(860 != -872)))))}
    v9 := []int{-9223372036854775808, -189, 9223372036854775807}
    var v10 bool = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int32 = 1
    var v2 string = "hello"
    v3 := make(map[int]int)
    var v4 bool = false
    v5 := make(map[string]int)
    var v6 uint8 = 127
    for fi11 := 0; fi11 < 14; fi11++ {
        v2 = v2
        v1 = v1
        _ = fi11
    }
    if (((1e10 + ((0.1 * (0.0 / 3.14159)) + ((0.0 * 1e10) - float64((float32(3.14159) - 1e10))))) <= (1e10 - ((float64(-2.25) * 0.0) / 0.1))) && v4) {
        if false {
            var v12 int8 = 125
            _ = v12
        }
        v2 = v2
        var v13 uint64 = 4294967295
        _ = v13
    } else {
        var v14 uint = 109
        var v15 int64 = 9223372036854775807
        var v16 float64 = 3.14159
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint32 = 110
        v19 := []int{}
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int16 = 1
    _ = len(v5)
    v20 += v20
    v2 = ("hello" + v2)
    var v21 int16 = 1000
    var v22 int16 = 834
    var v23 uint = 0
    for fi24 := 0; fi24 < 14; fi24++ {
        delete(v3, 1)
        v20 += (v22 & int16(v23))
        _ = fi24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        v1 = (v1 / (((v1 - (^(v1)))) | 1))
        v20 = v22
        var v26 int16 = v21
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
