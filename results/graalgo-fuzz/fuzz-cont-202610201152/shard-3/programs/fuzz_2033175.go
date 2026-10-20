package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint; F1 int16 }
func genFn0() uint16 {
    if false {
        for fi4 := 0; fi4 < 5; fi4++ {
            for fi5 := 0; fi5 < 13; fi5++ {
                fmt.Println(fi4, fi5)
                v6 := []int{}
                var v7 uint = 0
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = fi4
        }
    }
    v8 := []int{127, 128, 935}
    if true {
        var v9 uint16 = 39
        v9 = v9
        _ = v9
    } else {
        switch "foo bar" {
        case "\n\t":
            var v10 uint32 = 0
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            if (!false) {
                v8 = append(v8, (-(fi11)))
            }
            _ = fi11
        }
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        if (-32768 < -32768) {
            _ = len(v8)
            var v13 uint8 = 255
            _ = v13
        }
        _ = fi12
    }
    var v14 uint32 = 574
    _ = v14
    _ = v8
    return 1
}

func main() {
    var v1 string = ("foo bar" + "a")
    v2 := make(map[uint32]int)
    var v3 uint8 = 2
    fmt.Println(v3, v1)
    v15 := []int{}
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    v15 = append(v15, 2147483647)
    var v18 int = -409
    if false {
        var v19 int8 = 73
        _ = v19
    } else {
        _ = len(v15)
        var v20 uint32 = 2147483648
        _ = v20
    }
    var v21 bool = (-760 > 1)
    v1 = v1
    v22 := GenS1{F0: -1, F1: 936}
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v18, v3, fi23)
        v15 = append(v15, (fi23 ^ v18))
        _ = fi23
    }
    v18 = (-(v18))
    if v21 {
        switch (v18 | (v18 + 241)) {
        case 4294967295:
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 uint32 = 2147483647
                _ = fi24
                _ = v25
            }
        case -931:
            for fi26 := 0; fi26 < 17; fi26++ {
                v1 = v1
                switch -988 {
                case -1:
                    fmt.Println(v3)
                default:
                    v3 -= v3
                }
                v27 := GenS1{F0: 164, F1: 2147483647}
                _ = fi26
                _ = v27
            }
        case -9223372036854775808:
            v21 = false
        default:
            if v21 {
                if len(v15) > 0 {
                    v15[0] = v18
                }
                for fi28 := 0; fi28 < 9; fi28++ {
                    v15 = append(v15, 42)
                    switch v1 {
                    case "":
                        for fi29 := 0; fi29 < 12; fi29++ {
                            if v21 {
                                delete(v2, 2)
                                fmt.Println(v18)
                            }
                            _ = fi29
                        }
                    }
                    v3 = v3
                    _ = fi28
                }
            }
        }
        fmt.Println(v18)
    } else {
        v30 := GenS1{F0: -689, F1: 91}
        for fi31 := 0; fi31 < 2; fi31++ {
            for fi32 := 0; fi32 < 20; fi32++ {
                v33 := GenS2{F0: 1, F1: 42}
                if v21 {
                    switch (fi32 | (v18 & 32768)) {
                    case -780:
                        v1 = v1
                    case 32768:
                        v3 = 41
                    case 128:
                        v3 -= (v3 - (v3 * 255))
                    }
                }
                _ = fi32
                _ = v33
            }
            if (1000 != 1) {
                var v34 uint8 = v3
                for fi35 := 0; fi35 < 9; fi35++ {
                    var v36 uint16 = 0
                    _ = fi35
                    _ = v36
                }
                _ = v34
            } else {
                for fi37 := 0; fi37 < 16; fi37++ {
                    var v38 uint64 = 18446744073709551615
                    _ = fi37
                    _ = v38
                }
                switch "a" {
                case "\n\t":
                    if (1.5 < 3.14159) {
                        v39 := GenS2{F0: 65535, F1: 77}
                        for fi40 := 0; fi40 < 5; fi40++ {
                            var v41 bool = (!v21)
                            v3 = 127
                            _ = fi40
                            _ = v41
                        }
                        _ = v39
                    }
                case "":
                    if v21 {
                        v15 = append(v15, (fi31 | (v18 >> ((1) & 63))))
                    }
                case "foo bar":
                    for fi42 := 0; fi42 < 6; fi42++ {
                        var v43 int16 = -871
                        _ = fi42
                        _ = v43
                    }
                }
                v21 = v21
            }
            switch (v18 - 32767) {
            case 42:
                v15 = append(v15, fi31)
            case 345:
                fmt.Println(fi31)
            case 2147483647:
                v1 = ((v1 + ((("foo bar" + (v1 + v1)) + "foo bar") + v1)) + v1)
            }
            _ = fi31
        }
        _ = v30
    }
    var v44 int32 = 1
    v45 := GenS2{F0: 757, F1: 32767}
    v21 = ((v21 && (v21 && false)) && true)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v44
    _ = v45
}
