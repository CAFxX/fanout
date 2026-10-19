package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32) int8 {
    var v4 bool = (true || false)
    for fi5 := 0; fi5 < 12; fi5++ {
        if (!(42 <= 42)) {
            fmt.Println(p0)
        } else {
            if (v4 && (0 < 633)) {
                var v6 string = (("foo bar" + "\n\t") + "\n\t")
                if (18446744073709551615 <= 42) {
                    switch v6 {
                    case "a":
                        var v7 int8 = -128
                        _ = v7
                    case "\n\t":
                        v4 = (832 <= 754)
                    case "":
                        v8 := []int{(fi5 + fi5)}
                        _ = v8
                    default:
                        for fi9 := 0; fi9 < 10; fi9++ {
                            v10 := []int{fi5}
                            v4 = v4
                            switch v6 {
                            case "\n\t":
                                fmt.Println(fi5, p0, p0)
                                if ((142 > 2) || ((!false) || false)) {
                                    v4 = v4
                                } else {
                                    v11 := GenS1{F0: ""}
                                    var v12 uint = 42
                                    _ = v11
                                    _ = v12
                                }
                                v13 := GenS1{F0: (v6 + (("\n\t" + "a") + (("foo bar" + v6) + "a")))}
                                v14 := GenS1{F0: (v6 + (("" + (v6 + v6)) + v6))}
                                _ = v13
                                _ = v14
                            }
                            _ = fi9
                            _ = v10
                        }
                    }
                }
                _ = v6
            } else {
                var v15 int16 = 1000
                _ = v15
            }
            p0 -= p0
        }
        if (1 >= 42) {
            v16 := []int{fi5, fi5}
            var v17 uint64 = 131
            for fi18 := 0; fi18 < 20; fi18++ {
                switch fi18 {
                case 32768:
                    var v19 float64 = 0.0
                    _ = v19
                }
                var v20 int8 = -51
                _ = fi18
                _ = v20
            }
            _ = v16
            _ = v17
        } else {
            var v21 bool = ((!true) || v4)
            _ = v21
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 int16, p1 uint16) uint8 {
    var v22 int8 = 127
    if false {
        switch 4294967295 {
        case 348:
            var v23 int32 = 2147483647
            _ = v23
        case 128:
            p0 = 816
        }
        v22 += v22
    }
    _ = p0
    _ = p1
    _ = v22
    return 42
}

func main() {
    var v1 int = -223
    var v2 bool = (-2147483648 >= 42)
    var v3 uint64 = 42
    var v24 uint = 604
    if true {
        v3 = (v3 % (((v3 - 4294967295)) | 1))
    }
    var v25 int16 = 47
    v1 = (v1 * v1)
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := []int{}
        v28 := GenS2{F0: v25}
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v3, v1)
    var v29 int8 = 121
    v2 = (!v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
}
