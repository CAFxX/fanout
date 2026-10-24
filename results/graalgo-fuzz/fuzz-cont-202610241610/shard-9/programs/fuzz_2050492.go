package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int32 }

func main() {
    var v1 int16 = 0
    var v2 string = "a"
    var v3 int64 = -1
    var v4 string = v2
    var v5 int16 = 32767
    v6 := []uint32{100}
    v7 := GenS1{F0: (v5 >> ((992) & 63)), F1: v3, F2: 42}
    v5 = v5
    v3 = (v3 | v3)
    var v8 int64 = v3
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v4, v3)
        _ = fi9
    }
    var v10 uint = 65535
    if (!(false && false)) {
        if (false && true) {
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 bool = (!(1e10 <= (0.0 - -2.25)))
                fmt.Println(v4, v1, v2)
                if v12 {
                    switch (v2 + v2) {
                    case "a":
                        switch v2 {
                        case "":
                            _ = len(v6)
                        case "a":
                            fmt.Println(v1, fi11)
                        }
                    case "":
                        v12 = (!v12)
                    case "hello":
                        var v13 uint8 = 255
                        _ = v13
                    }
                    v10 -= v10
                }
                _ = fi11
                _ = v12
            }
            v8 = (v8 >> ((v10) & 63))
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(v10)
                _ = fi14
            }
        }
        if (42 != 2) {
            var v15 float64 = (((3.14159 - 0.1) - (0.0 / -2.25)) * ((1.5 * 1.5) - 3.14159))
            var v16 bool = false
            _ = v15
            _ = v16
        } else {
            fmt.Println(v4, v10, v1)
        }
    }
    fmt.Println(v1, v2)
    v3 -= v3
    if true {
        switch 32768 {
        case 32767:
            var v17 uint8 = 127
            _ = v17
        }
        var v18 int = 42
        var v19 uint64 = 1
        _ = v18
        _ = v19
    }
    fmt.Println(v4, v4, v5)
    var v20 int = 0
    var v21 int64 = v8
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := GenS1{F0: (v5 - 459), F1: v8, F2: 0}
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: v5, F1: -783, F2: 42}
    v5 = v5
    v5 = (v5 % (((^(v1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
