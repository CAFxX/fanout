package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }
func genFn0() uint8 {
    var v6 int = 980
    switch ("\n\t" + ("a" + "foo bar")) {
    case "\n\t":
        if (2147483647 >= 2) {
            fmt.Println(v6)
        }
    }
    if true {
        var v7 uint8 = 1
        var v8 uint8 = 122
        for fi9 := 0; fi9 < 4; fi9++ {
            switch v6 {
            case 542:
                v8 += (v8 << ((4294967295) & 63))
            case 9223372036854775807:
                var v10 uint64 = 4294967295
                _ = v10
            default:
                v11 := GenS1{F0: -728}
                _ = v11
            }
            var v12 float32 = -2.25
            v12 += 1.5
            _ = fi9
            _ = v12
        }
        _ = v7
        _ = v8
    }
    if ((132 < 18446744073709551615) && ((float64(0.0) + 0.1) > 0.0)) {
        fmt.Println(v6)
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS2{F0: 128, F1: 2, F2: 1}
            fmt.Println(v6, v6)
            if (!false) {
                if (759 < 9223372036854775807) {
                    if (0 != 321) {
                        if (false || ((("a" + ("\n\t" + "a")) > ("hello" + "")) || false)) {
                            var v15 float32 = -2.25
                            fmt.Println(fi13, v6, v15)
                            var v16 uint = 850
                            _ = v15
                            _ = v16
                        }
                        switch "a" {
                        case "":
                            var v17 uint16 = 156
                            _ = v17
                        }
                    } else {
                        var v18 int16 = 1000
                        var v19 int8 = 0
                        _ = v18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 13; fi20++ {
                        if true {
                            var v21 int8 = 0
                            _ = v21
                        } else {
                            var v22 string = "a"
                            _ = v22
                        }
                        _ = fi20
                    }
                    if (true && false) {
                        v23 := []int{fi13}
                        _ = v23
                    }
                }
            }
            _ = fi13
            _ = v14
        }
    } else {
        var v24 bool = (85 != 18446744073709551615)
        _ = v24
    }
    _ = v6
    return 137
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -551
    var v3 float64 = 0.0
    var v4 int8 = -111
    var v5 bool = false
    var v25 int64 = -868
    var v26 bool = (!v5)
    if (1.5 < 0.1) {
        var v27 int64 = (v25 - v25)
        v28 := []int{(-(v2)), (v2 * -2147483648)}
        switch ("hello" + "a") {
        case "\n\t":
            switch v2 {
            case -2147483648:
                var v29 float32 = (0.0 * 1e10)
                _ = v29
            case 2147483647:
                var v30 int64 = (v27 - int64(v2))
                _ = v30
            }
        }
        _ = v27
        _ = v28
    } else {
        var v31 uint32 = 2
        _ = v31
    }
    var v32 bool = (!v5)
    v5 = true
    v33 := GenS1{F0: -11}
    v25 += 9223372036854775807
    var v34 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
