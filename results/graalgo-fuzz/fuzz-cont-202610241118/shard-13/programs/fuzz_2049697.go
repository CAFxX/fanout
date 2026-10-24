package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int32 = 421
    var v2 uint = 0
    var v3 string = "hello"
    var v4 string = (v3 + "")
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{fi5, 42, fi5, (fi5 & -474)}
        if ((3.14159 < -2.25) || false) {
            v7 := GenS1{F0: (fi5 * (fi5 % (((fi5 + fi5)) | 1)))}
            _ = v7
        }
        var v8 bool = true
        _ = fi5
        _ = v6
        _ = v8
    }
    if true {
        fmt.Println(v3)
        var v9 string = v3
        v3 = (v4 + "foo bar")
        _ = v9
    }
    v10 := []int{-2147483648, 127}
    if ("hello" >= ("hello" + v4)) {
        v1 = -1
    }
    var v11 uint32 = 4294967295
    v2 += v2
    if (!(true || (32768 == 32767))) {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 bool = (!false)
            switch (fi12 * fi12) {
            case 0:
                for fi14 := 0; fi14 < 17; fi14++ {
                    fmt.Println(v1)
                    switch v3 {
                    case "foo bar":
                        var v15 string = (v4 + v4)
                        _ = v15
                    }
                    if len(v10) > 0 {
                        v10[0] = fi12
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    if len(v10) > 0 {
                        v10[0] = (fi14 | fi12)
                    }
                    _ = fi14
                }
            }
            _ = len(v10)
            _ = fi12
            _ = v13
        }
    }
    if ((2 <= 1) || (!(!true))) {
        var v16 string = "\n\t"
        switch 576 {
        case 1:
            var v17 int8 = 1
            _ = v17
        default:
            if (false || false) {
                v18 := GenS1{F0: 1}
                var v19 int32 = 696
                _ = v18
                _ = v19
            } else {
                if false {
                    var v20 uint32 = 65535
                    if false {
                        if len(v10) > 0 {
                            v10[0] = -7
                        }
                        if ((!(32767 < 128)) && (!((v1 / (((v1 / ((v1) | 1))) | 1)) == (v1 + v1)))) {
                            v21 := GenS1{F0: 128}
                            var v22 int = 32768
                            _ = v21
                            _ = v22
                        }
                        v23 := GenS1{F0: int(v1)}
                        _ = v23
                    } else {
                        var v24 uint32 = v20
                        var v25 uint32 = v20
                        var v26 int8 = 0
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    _ = v20
                }
                var v27 uint64 = 0
                var v28 int64 = -399
                _ = v27
                _ = v28
            }
        }
        _ = v16
    }
    var v29 uint16 = 32767
    if true {
        var v30 float64 = 1.5
        fmt.Println(v30)
        _ = v30
    }
    fmt.Println(v29, v2, v3)
    v1 += -390
    var v31 uint64 = 2
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v4
}
