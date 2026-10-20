package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint32 }

func main() {
    var v1 uint32 = 965
    var v2 uint64 = 18446744073709551615
    var v3 float64 = 3.14159
    var v4 int64 = 9223372036854775807
    var v5 uint16 = 274
    v6 := make(map[string]bool)
    if (42 == -571) {
        v7 := []int{-836}
        _ = v7
    }
    if (false && (32767 <= 32767)) {
        var v8 uint = 265
        fmt.Println(v1, v3, v8)
        v9 := []int{4294967295, -9223372036854775808}
        _ = v8
        _ = v9
    }
    var v10 bool = true
    v3 += v3
    v11 := []int{127, 55}
    fmt.Println(v2, v5, v3)
    if (0.1 >= (v3 - v3)) {
        var v12 float64 = 1.5
        _ = v12
    } else {
        var v13 uint = 42
        var v14 uint16 = v5
        fmt.Println(v4)
        _ = v13
        _ = v14
    }
    if (558 != 65535) {
        if len(v11) > 0 {
            v11[0] = 222
        }
        v11 = append(v11, 662)
        v15 := GenS1{F0: (^(v1)), F1: v4}
        _ = v15
    }
    if (false || v10) {
        v16 := v6[(((("\n\t" + ("" + "foo bar")) + "foo bar") + ("hello" + "")) + "foo bar")]
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 bool = (v10 && (105 != 0))
        switch (^(fi17)) {
        case -624:
            v2 = 1
        case 128:
            var v19 int8 = -1
            _ = v19
        case -601:
            if v18 {
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 string = "foo bar"
                    var v22 uint8 = 91
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        default:
            var v23 int32 = 284
            _ = v23
        }
        if true {
            fmt.Println(v1, v2)
            var v24 uint64 = (v2 >> ((4294967295) & 63))
            v4 -= v4
            _ = v24
        } else {
            for fi25 := 0; fi25 < 1; fi25++ {
                v10 = (!(!((v1 & (v1 >> ((985) & 63))) != 547)))
                _ = fi25
            }
            fmt.Println(v18)
            var v26 int32 = -1
            _ = v26
        }
        _ = fi17
        _ = v18
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 int64 = v4
        _ = fi27
        _ = v28
    }
    var v29 int8 = 1
    var v30 int = 127
    _ = len(v6)
    if false {
        if v10 {
            v10 = v10
            if v10 {
                fmt.Println(v10, v10, v1)
            } else {
                var v31 int64 = v4
                fmt.Println(v4, v2)
                v32 := GenS1{F0: (v1 | (v1 - v1)), F1: v31}
                _ = v31
                _ = v32
            }
        }
        if false {
            for fi33 := 0; fi33 < 6; fi33++ {
                if v10 {
                    v29 = -15
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v35 := GenS1{F0: v1, F1: v4}
                        _ = len(v6)
                        var v36 uint8 = 2
                        _ = len(v11)
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
                var v37 float64 = v3
                if len(v11) > 0 {
                    v11[0] = (^(v30))
                }
                _ = fi33
                _ = v37
            }
        } else {
            var v38 int32 = -452
            switch 32767 {
            case -657:
                v30 = v30
            case -2147483648:
                fmt.Println(v10, v5)
            case 1:
                delete(v6, "\n\t")
            }
            v39 := GenS2{F0: v29, F1: 39, F2: 736}
            _ = v38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
