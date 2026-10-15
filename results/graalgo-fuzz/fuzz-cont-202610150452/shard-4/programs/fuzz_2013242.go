package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint64) int32 {
    p0 += (p0 | 2)
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 float64 = 1.5
        if ((801 > 9223372036854775807) || (!false)) {
            v8 := []int{fi6}
            _ = v8
        } else {
            var v9 int = (fi6 * 900)
            var v10 uint8 = 21
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 int32 = 2147483647
                var v13 uint32 = 65535
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
            _ = v9
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            if true {
                if false {
                    v15 := GenS1{F0: 0}
                    var v16 int32 = 42
                    var v17 uint = 1
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS1{F0: 113}
                    _ = v18
                }
                switch fi14 {
                case 32767:
                    var v19 string = "a"
                    _ = v19
                default:
                    fmt.Println(v7, fi14, fi14)
                }
            }
            _ = fi14
        }
        _ = fi6
        _ = v7
    }
    var v20 float64 = 1e10
    v21 := GenS1{F0: -906}
    _ = p0
    _ = v20
    _ = v21
    return -244
}
func genFn1(p0 int, p1 uint) int8 {
    var v22 float64 = -2.25
    v23 := []int{}
    if (13 <= 127) {
        v24 := GenS2{F0: 255}
        v23 = append(v23, 128)
        _ = v24
    } else {
        switch "" {
        case "foo bar":
            fmt.Println(p0)
        case "hello":
            fmt.Println(p0, v22)
        case "\n\t":
            fmt.Println(p1, p0, p0)
        }
        v22 = v22
        var v25 string = (("hello" + (("foo bar" + ("" + ("hello" + "\n\t"))) + "")) + "")
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 127
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    var v3 int32 = -482
    var v4 int16 = -528
    v5 := make(map[string]bool)
    for fi26 := 0; fi26 < 10; fi26++ {
        if (0 >= (fi26 * -2147483648)) {
            v4 += v4
            fmt.Println(v3, v3)
            _ = len(v1)
        } else {
            fmt.Println(v3)
            v27 := GenS3{F0: genFn1((fi26 - (fi26 & fi26)), 18446744073709551615)}
            var v28 string = ("\n\t" + ("hello" + (("\n\t" + "\n\t") + "foo bar")))
            _ = v27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println(v3)
    if true {
        for fi29 := 0; fi29 < 11; fi29++ {
            switch "\n\t" {
            case "\n\t":
                for fi30 := 0; fi30 < 17; fi30++ {
                    v31 := []int{fi30, fi30, (fi29 | fi30), (fi30 / (((^(fi30))) | 1))}
                    _ = fi30
                    _ = v31
                }
            case "a":
                if (!(true && ((3.14159 * (-2.25 / 0.1)) <= 1e10))) {
                    for fi32 := 0; fi32 < 1; fi32++ {
                        v33 := GenS2{F0: 2}
                        fmt.Println(fi29, v3, fi32)
                        v4 = 276
                        _ = fi32
                        _ = v33
                    }
                }
            case "foo bar":
                if false {
                    switch fi29 {
                    case 2147483647:
                        for fi34 := 0; fi34 < 12; fi34++ {
                            v35 := GenS4{F0: -511}
                            v4 += v4
                            _ = fi34
                            _ = v35
                        }
                    }
                    switch "" {
                    case "a":
                        fmt.Println(v4, v3)
                    default:
                        v4 += (v4 << ((356) & 63))
                    }
                } else {
                    v3 += (v3 << ((18446744073709551615) & 63))
                    v36 := GenS5{F0: (v3 | genFn0(4294967295)), F1: (v3 % (((^(v3))) | 1))}
                    switch "foo bar" {
                    case "foo bar":
                        v4 += (v4 % (((v4 >> ((255) & 63))) | 1))
                    case "\n\t":
                        switch "" {
                        case "foo bar":
                            for fi37 := 0; fi37 < 20; fi37++ {
                                fmt.Println(v3)
                                v38 := []int{(fi37 | fi29), fi37}
                                _ = fi37
                                _ = v38
                            }
                        case "hello":
                            fmt.Println(fi29)
                        case "":
                            if (v3 >= (v3 & (v3 >> ((989) & 63)))) {
                                var v39 uint = 69
                                v40 := GenS3{F0: 100}
                                _ = v39
                                _ = v40
                            } else {
                                var v41 int = 2147483647
                                fmt.Println(v41, v3)
                                _ = len(v5)
                                _ = v41
                            }
                        }
                    case "a":
                        switch (fi29 - -657) {
                        case 0:
                            v42 := v5["hello"]
                            _ = v42
                        }
                    default:
                        for fi43 := 0; fi43 < 5; fi43++ {
                            var v44 int = 1
                            var v45 uint8 = 42
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    }
                    _ = v36
                }
            }
            for fi46 := 0; fi46 < 2; fi46++ {
                v3 -= genFn0(628)
                _ = fi46
            }
            var v47 uint8 = 42
            _ = fi29
            _ = v47
        }
        v48 := []int{973, 2147483647, 4294967295}
        v4 -= (v4 >> ((4294967295) & 63))
        _ = v48
    } else {
        var v49 bool = true
        v4 += v4
        _ = v49
    }
    for fi50 := 0; fi50 < 13; fi50++ {
        var v51 float64 = 0.1
        var v52 int64 = -18
        _ = fi50
        _ = v51
        _ = v52
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        v4 -= 32767
        fmt.Println(fi53, v3)
        _ = fi53
    }
    var v54 uint64 = 0
    if false {
        v55 := GenS6{F0: 183}
        v56 := []int{4294967295, 0}
        _ = len(v56)
        _ = v55
        _ = v56
    }
    _ = len(v5)
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
}
