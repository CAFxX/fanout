package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 int8 = 100
    var v4 int64 = -24
    v4 -= (v4 >> ((0) & 63))
    var v5 float64 = ((1.5 + (float64(0.0) / 0.1)) - 1e10)
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v1["a"] = ("a" + ("\n\t" + ""))
            _ = fi7
        }
        _ = fi6
    }
    v4 -= v4
    if (1 >= 464) {
        switch (("hello" + "a") + "foo bar") {
        case "a":
            v8 := []int{32767, -1, 127, 32767}
            _ = v8
        default:
            fmt.Println(v5, v3, v4)
        }
        if ((true && (0 == 2)) || ("" <= ("hello" + "a"))) {
            var v9 float32 = 0.1
            _ = v9
        } else {
            fmt.Println(v4)
            v10 := []int{-232, 316, 534}
            _ = v10
        }
        switch -501 {
        case 4294967295:
            v2["hello"] = 288
        case 1:
            var v11 uint64 = 4294967295
            _ = v11
        case 32767:
            var v12 int8 = 78
            _ = v12
        }
    } else {
        if false {
            var v13 float32 = (0.1 / (0.1 / ((float32(1.5) / (float32(((v5 / v5) / (v5 / v5))) / 0.1)) * 1.5)))
            v14 := GenS1{F0: ("" + ("a" + ("a" + "a"))), F1: 0, F2: false}
            var v15 int16 = -32768
            _ = v13
            _ = v14
            _ = v15
        } else {
            switch 128 {
            case 128:
                var v16 int8 = -128
                _ = v16
            case 127:
                fmt.Println(v3, v5)
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                switch (fi17 << ((18446744073709551615) & 63)) {
                case 2147483647:
                    v18 := GenS1{F0: "a", F1: -32768, F2: false}
                    _ = v18
                case -337:
                    v3 += (-(v3))
                case -791:
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v20 := []int{fi17, fi19, fi19, fi19}
                        fmt.Println(v4, v4, v5)
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                        fmt.Println(v3)
                        var v21 int64 = (v4 << ((0) & 63))
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi17
            }
            var v22 string = ("" + "a")
            _ = v22
        }
        fmt.Println(v5, v4)
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        switch fi23 {
        case 1:
            v4 -= (^(v4))
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            if true {
                var v25 int16 = 531
                _ = v25
            }
            fmt.Println(fi23, v5)
            _ = fi24
        }
        v26 := []int{fi23, fi23}
        _ = fi23
        _ = v26
    }
    var v27 int = 4294967295
    if true {
        _ = len(v1)
        var v28 uint16 = 659
        _ = v28
    } else {
        _ = len(v1)
        v29 := []int{}
        v5 += (v5 / (-2.25 + (v5 * 1e10)))
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
