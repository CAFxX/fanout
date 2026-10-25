package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16) uint16 {
    p0 = -937
    v6 := []int{}
    _ = p0
    _ = v6
    return 727
}

func main() {
    var v1 float64 = -2.25
    var v2 string = ("hello" + "\n\t")
    v3 := make(map[uint32]int)
    var v4 uint32 = 25
    var v5 bool = true
    var v7 int = -136
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v2 = "hello"
            _ = fi9
        }
        _ = fi8
    }
    v2 = ((v2 + ((v2 + v2) + ("hello" + (v2 + v2)))) + v2)
    var v10 int64 = 129
    if v5 {
        if (2 == 1) {
            v2 = "\n\t"
            if (!(false || v5)) {
                var v11 uint16 = 344
                v5 = (uint16(v4) == 0)
                v12 := v3[(v4 - 612)]
                _ = v11
                _ = v12
            } else {
                fmt.Println(v1)
            }
            v13 := GenS1{F0: -128}
            _ = v13
        }
        v14 := GenS1{F0: -1}
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 int16 = 0
            switch (v7 >> ((4294967295) & 63)) {
            case 42:
                switch ("" + ("foo bar" + "foo bar")) {
                case "hello":
                    v17 := []int{(v7 + fi15), v7, -488, 127}
                    _ = v17
                default:
                    if (float64((((1e10 * (0.0 / 3.14159)) - ((3.14159 - 0.0) * 1.5)) + 3.14159)) > float64((0.1 - (3.14159 * (0.0 + 1e10))))) {
                        v18 := GenS1{F0: 1}
                        _ = v18
                    } else {
                        v4 = (v4 >> ((422) & 63))
                        var v19 int64 = v10
                        v20 := []int{fi15, 793, -1, 0}
                        _ = v19
                        _ = v20
                    }
                }
            case -41:
                switch v2 {
                case "a":
                    var v21 uint16 = 646
                    _ = v21
                }
            case 4294967295:
                v16 = (^(v16))
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        fmt.Println(v1)
    }
    var v22 bool = v5
    if (!(127 > 128)) {
        v23 := GenS1{F0: 11}
        _ = v23
    } else {
        v24 := []int{}
        var v25 uint16 = 32767
        fmt.Println(v7, v25, v2)
        _ = v24
        _ = v25
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v26 float32 = 0.1
        _ = v26
    case 42:
        var v27 uint = 65535
        _ = v27
    default:
        for fi28 := 0; fi28 < 17; fi28++ {
            fmt.Println(v10)
            _ = fi28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
