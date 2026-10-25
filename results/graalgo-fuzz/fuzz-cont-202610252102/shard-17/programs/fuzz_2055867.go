package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 bool = true
    v2 := []int{145, -322, -106, 1}
    v3 := []int{42}
    var v4 int32 = -929
    for fi5 := 0; fi5 < 13; fi5++ {
        v1 = (42 == 32767)
        switch (fi5 + -32) {
        case -991:
            v1 = v1
        }
        _ = fi5
    }
    var v6 string = "a"
    fmt.Println(v4)
    var v7 int32 = -2147483648
    var v8 float32 = 0.1
    var v9 int16 = 1000
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 string = "hello"
        v4 = (v4 & (-(v7)))
        v7 -= -1
        _ = fi10
        _ = v11
    }
    fmt.Println(v6, v6, v4)
    v1 = (0.0 >= (((1e10 - 1.5) + 3.14159) - ((0.0 + 3.14159) * 0.1)))
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 uint32 = 594
        if v1 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v1)
            switch fi12 {
            case 317:
                var v14 uint32 = 2147483647
                _ = v14
            case 32768:
                var v15 int64 = -636
                _ = v15
            case 1:
                v6 = v6
            default:
                for fi16 := 0; fi16 < 7; fi16++ {
                    v7 = int32(v9)
                    _ = fi16
                }
            }
        }
        fmt.Println(v7, v13, v4)
        _ = fi12
        _ = v13
    }
    var v17 uint16 = 723
    var v18 float32 = v8
    fmt.Println(v4)
    if v1 {
        if true {
            var v19 string = ((v6 + (("foo bar" + v6) + (v6 + ("a" + "a")))) + v6)
            if v1 {
                v7 = -991
                var v20 uint64 = 2
                v21 := GenS1{F0: 42}
                _ = v20
                _ = v21
            } else {
                var v22 int32 = 0
                v9 -= (v9 * -998)
                _ = v22
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v19
        }
    } else {
        if (v1 || v1) {
            v23 := GenS1{F0: -77}
            _ = v23
        } else {
            var v24 uint32 = 76
            _ = v24
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            fmt.Println(v7, fi25)
            var v26 bool = v1
            _ = fi25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v4 -= v4
        _ = fi27
    }
    var v28 bool = (!(2 > 0))
    var v29 int32 = 356
    v30 := GenS1{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
