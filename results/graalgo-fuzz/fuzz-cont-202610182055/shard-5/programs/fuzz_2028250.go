package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0() int16 {
    v6 := []int{-9223372036854775808}
    if (1e10 == ((((1e10 / -2.25) + (float64(float32(float64(((float32(-2.25) - 0.1) * -2.25)))) * (-2.25 + 1e10))) + 0.0) - 0.1)) {
        v6 = append(v6, 4294967295)
        if false {
            _ = len(v6)
            var v7 int64 = 9223372036854775807
            var v8 uint64 = 18446744073709551615
            _ = v7
            _ = v8
        }
    } else {
        _ = len(v6)
        var v9 float32 = 1.5
        _ = v9
    }
    if true {
        var v10 int8 = 110
        var v11 uint64 = 42
        v12 := GenS1{F0: 2, F1: true, F2: 0}
        _ = v10
        _ = v11
        _ = v12
    } else {
        switch "a" {
        case "hello":
            var v13 int8 = -1
            _ = v13
        case "\n\t":
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS1{F0: 599, F1: (42 != 955), F2: 609}
                var v16 float32 = float32((3.14159 - 1.5))
                _ = fi14
                _ = v15
                _ = v16
            }
        case "":
            var v17 bool = (!true)
            _ = v17
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v18 uint32 = 229
    v19 := GenS2{F0: 1000, F1: 1}
    _ = v18
    _ = v19
    _ = v6
    return 1000
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 772
    var v3 int32 = -1
    v4 := make(map[int]uint32)
    var v5 bool = (!("foo bar" > ("\n\t" + (("hello" + "") + "foo bar"))))
    var v20 int = 0
    var v21 float32 = -2.25
    v4[2147483647] = 65535
    fmt.Println(v5)
    switch ("\n\t" + "a") {
    case "hello":
        var v22 bool = ("a" > "hello")
        _ = v22
    }
    var v23 string = ("hello" + "")
    if ((v5 && true) && v5) {
        var v24 int32 = -110
        v25 := []int{}
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 uint32 = 2147483647
            _ = fi26
            _ = v27
        }
    }
    v21 = 3.14159
    v28 := GenS1{F0: 65535, F1: true, F2: 841}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
