package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint) string {
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(p1, fi5)
        _ = fi5
    }
    v6 := []int{-972, -921, -108}
    var v7 uint32 = (p0 << ((p1) & 63))
    if true {
        v6 = append(v6, 218)
        var v8 int = -2147483648
        _ = v8
    }
    if len(v6) > 0 {
        v6[0] = 127
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ("" + "foo bar")
}

func main() {
    var v1 uint = 2
    var v2 int64 = 0
    var v3 uint = 255
    var v4 bool = (65535 >= 918)
    v9 := []int{42, -1}
    v10 := GenS1{F0: 0}
    _ = len(v9)
    var v11 float64 = 1e10
    if v4 {
        switch "" {
        case "a":
            var v12 bool = false
            _ = v12
        case "foo bar":
            if v4 {
                _ = len(v9)
                var v13 uint64 = 0
                _ = v13
            }
        case "\n\t":
            v2 = 42
        }
        var v14 int16 = -238
        v1 = 42
        _ = v14
    }
    var v15 int16 = -1
    var v16 float32 = 0.1
    var v17 uint32 = 263
    for fi18 := 0; fi18 < 10; fi18++ {
        _ = len(v9)
        if false {
            switch fi18 {
            case 4294967295:
                v2 = v2
            case -1:
                switch (fi18 / ((9223372036854775807) | 1)) {
                case 618:
                    if true {
                        var v19 int16 = (-(v15))
                        fmt.Println(fi18, v15)
                        _ = v19
                    }
                }
            }
        } else {
            for fi20 := 0; fi20 < 10; fi20++ {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                var v21 uint = (v3 >> (((v3 + (v3 & (^(v3))))) & 63))
                fmt.Println(v3, v15)
                _ = fi20
                _ = v21
            }
        }
        _ = fi18
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        if (((-(fi22)) >= fi22) && (!v4)) {
            for fi23 := 0; fi23 < 12; fi23++ {
                switch (-(fi22)) {
                case -316:
                    switch ("\n\t" + ("" + "\n\t")) {
                    case "\n\t":
                        if len(v9) > 0 {
                            v9[0] = (fi23 - fi22)
                        }
                    default:
                        fmt.Println(v2, v17)
                    }
                default:
                    var v24 int = -777
                    _ = v24
                }
                if v4 {
                    _ = len(v9)
                    if v4 {
                        var v25 float32 = v16
                        _ = len(v9)
                        v4 = (v4 && (v4 && (!v4)))
                        _ = v25
                    }
                }
                v3 += v1
                _ = fi23
            }
        } else {
            for fi26 := 0; fi26 < 15; fi26++ {
                v11 += ((1.5 - ((v11 - (1e10 / v11)) - (float64(3.14159) + (v11 / v11)))) * -2.25)
                v4 = true
                _ = fi26
            }
            var v27 int64 = v2
            switch fi22 {
            case 4294967295:
                var v28 uint = (v1 ^ 4294967295)
                _ = v28
            }
            _ = v27
        }
        v15 = v15
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
