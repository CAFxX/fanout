package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    var v1 string = ("\n\t" + ("\n\t" + "\n\t"))
    var v2 float32 = 0.0
    var v3 int32 = -2147483648
    if (187 != 793) {
        var v4 int8 = -128
        v5 := []int{961, 197}
        _ = v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := []int{(fi6 & (^(fi6))), fi6, fi6}
            switch fi6 {
            case -1:
                v7 = append(v7, (^(fi6)))
            case 4294967295:
                switch v1 {
                case "\n\t":
                    v8 := GenS1{F0: 1}
                    _ = v8
                default:
                    v1 = v1
                }
            default:
                for fi9 := 0; fi9 < 20; fi9++ {
                    if (!true) {
                        if false {
                            switch (v1 + "foo bar") {
                            case "hello":
                                fmt.Println(fi6, v3, fi9)
                            case "a":
                                var v10 uint64 = 2
                                _ = v10
                            }
                        }
                        v7 = append(v7, fi9)
                    }
                    _ = fi9
                }
            }
            _ = fi6
            _ = v7
        }
        fmt.Println(v2, v2, v1)
    }
    v11 := GenS1{F0: 18446744073709551615}
    var v12 int8 = 100
    if (v3 <= (v3 - (v3 % ((2147483647) | 1)))) {
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 int8 = (^(v12))
            v3 = 462
            v15 := GenS2{F0: fi13, F1: (v3 << ((1) & 63))}
            _ = fi13
            _ = v14
            _ = v15
        }
        if false {
            if false {
                switch ((v1 + (v1 + v1)) + "\n\t") {
                case "foo bar":
                    var v16 int16 = 1000
                    _ = v16
                case "\n\t":
                    var v17 string = v1
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v2 += ((float32(-2.25) * 0.0) - (v2 - (v2 * ((v2 / -2.25) * ((v2 + v2) + ((1.5 - 0.1) + v2))))))
                        _ = fi18
                    }
                }
                var v19 uint16 = 262
                v20 := []int{-195}
                _ = v19
                _ = v20
            }
            v21 := []int{-1, 32768}
            var v22 bool = (0 >= 550)
            _ = v21
            _ = v22
        }
    } else {
        switch 127 {
        case 128:
            v23 := []int{804, -37, -9223372036854775808, 640}
            _ = v23
        default:
            var v24 string = v1
            _ = v24
        }
        v25 := []int{2147483647, -1, -741}
        _ = v25
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        for fi27 := 0; fi27 < 9; fi27++ {
            if (365 != 1) {
                v2 -= 1.5
                v28 := []int{(fi26 << ((255) & 63)), (fi27 & fi27)}
                v29 := GenS1{F0: 0}
                _ = v28
                _ = v29
            }
            _ = fi27
        }
        switch "" {
        case "foo bar":
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 uint16 = 32767
                v32 := GenS2{F0: fi26, F1: (v3 | -670)}
                _ = fi30
                _ = v31
                _ = v32
            }
        default:
            for fi33 := 0; fi33 < 12; fi33++ {
                switch fi26 {
                case 127:
                    var v34 uint8 = 0
                    _ = v34
                case 9223372036854775807:
                    v35 := []int{-1, 32768, fi26, fi26}
                    _ = v35
                case 42:
                    v36 := GenS1{F0: 113}
                    _ = v36
                default:
                    var v37 int64 = 1
                    if false {
                        v38 := GenS1{F0: 51}
                        _ = v38
                    }
                    v1 = ""
                    var v39 int32 = -836
                    _ = v37
                    _ = v39
                }
                _ = fi33
            }
        }
        var v40 bool = (2147483647 == v3)
        _ = fi26
        _ = v40
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        v2 = (float32(1e10) * (1.5 - 0.1))
        _ = fi41
    }
    v2 += (-2.25 * v2)
    if true {
        for fi42 := 0; fi42 < 12; fi42++ {
            for fi43 := 0; fi43 < 10; fi43++ {
                v44 := []int{fi43, fi43, (fi43 + (fi43 - 9223372036854775807)), 4294967295}
                v45 := GenS2{F0: fi42, F1: v3}
                _ = fi43
                _ = v44
                _ = v45
            }
            _ = fi42
        }
    } else {
        v3 -= (v3 + (v3 & v3))
        v46 := []int{-2147483648, 32767}
        _ = v46
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        var v48 uint32 = 2
        _ = fi47
        _ = v48
    }
    var v49 uint8 = 42
    v50 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v49
    _ = v50
}
