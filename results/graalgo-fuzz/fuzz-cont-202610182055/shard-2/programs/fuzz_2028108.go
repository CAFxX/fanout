package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int64 }
func genFn0() int8 {
    if true {
        var v5 uint8 = 127
        var v6 int8 = 89
        if (0.0 >= 0.1) {
            v6 -= (v6 ^ -128)
            var v7 int16 = -1
            _ = v7
        }
        _ = v5
        _ = v6
    } else {
        var v8 int64 = 42
        _ = v8
    }
    var v9 int = -664
    v9 = v9
    var v10 bool = false
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v9 += fi11
            var v13 float32 = float32(float64((1.5 + ((0.0 / ((-2.25 * 0.0) - (1e10 - 3.14159))) - float32(1.5)))))
            v14 := GenS1{F0: -1, F1: 2147483647, F2: -463}
            v15 := []int{fi12, (v9 ^ fi11), 916}
            var v16 int = (fi11 % ((9223372036854775807) | 1))
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
            _ = v16
        }
        var v17 bool = ((42 != 187) || (!(false && true)))
        var v18 int16 = 1000
        _ = fi11
        _ = v17
        _ = v18
    }
    _ = v10
    _ = v9
    return 100
}
func genFn1() string {
    var v19 bool = false
    var v20 string = (("foo bar" + "hello") + (("" + (("foo bar" + ("foo bar" + (("\n\t" + "hello") + ("a" + "foo bar")))) + "a")) + "hello"))
    if ((-2.25 - 0.0) >= ((0.1 * 0.1) - 3.14159)) {
        var v21 uint = 65535
        _ = v21
    } else {
        v22 := []int{1, 1}
        _ = v22
    }
    v23 := []int{1, -253, 4294967295, 9223372036854775807}
    _ = v19
    _ = v20
    _ = v23
    return v20
}

func main() {
    v1 := []string{"a", "", ("" + "a")}
    v2 := []int{42, -552}
    v3 := []uint32{583, 2}
    v4 := make(map[int]string)
    var v24 uint16 = 255
    v4[4294967295] = ""
    fmt.Println(v24)
    var v25 string = "a"
    if true {
        fmt.Println(v25)
        _ = len(v4)
    }
    if false {
        if true {
            if true {
                fmt.Println(v25, v25)
            }
        }
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 uint8 = 1
            _ = len(v1)
            if true {
                var v29 uint = 42
                var v30 uint = v29
                v25 = "\n\t"
                _ = v29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi26
    }
    fmt.Println(v25, v24)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
