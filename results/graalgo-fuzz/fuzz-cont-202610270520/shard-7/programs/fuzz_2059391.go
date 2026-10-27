package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 bool = false
    var v2 float64 = -2.25
    var v3 uint8 = 127
    v1 = (-9223372036854775808 == 156)
    v4 := GenS1{F0: (^(v3))}
    v5 := GenS2{F0: -414}
    var v6 uint = 0
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 2; fi8++ {
        if (!(1 <= 1)) {
            v6 = v6
            var v9 float32 = (3.14159 / 1.0)
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{(fi8 << ((v6) & 63)), (-(fi8))}
            v12 := GenS2{F0: 1000}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    var v13 float64 = 0.1
    for fi14 := 0; fi14 < 3; fi14++ {
        if true {
            var v15 bool = (0 == -1)
            v2 += 0.1
            _ = v15
        }
        var v16 float64 = (v13 / ((v13 * (-2.25 / v13)) - 3.14159))
        _ = fi14
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: v3}
            v6 = (v6 & (v6 ^ (v6 * (^(v6)))))
            _ = fi18
            _ = v19
        }
        v3 = 37
        if ((true && (!(!((v3 % ((v3) | 1)) >= (^(v3)))))) && true) {
            for fi20 := 0; fi20 < 18; fi20++ {
                v21 := []int{fi20, fi20}
                _ = fi20
                _ = v21
            }
        }
        _ = fi17
    }
    fmt.Println(v6, v3, v2)
    v22 := []int{128, 0}
    if ((v1 && (!true)) && (((("foo bar" + "\n\t") + "hello") + "\n\t") == "a")) {
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 bool = v1
            switch (fi23 & (-(fi23))) {
            case 32768:
                if true {
                    for fi25 := 0; fi25 < 5; fi25++ {
                        v24 = (false && (-2147483648 != -8))
                        fmt.Println(v24, v13)
                        if len(v22) > 0 {
                            v22[0] = fi23
                        }
                        _ = fi25
                    }
                    v24 = true
                    v13 += (v13 / 3.14159)
                } else {
                    fmt.Println(v6, v13, v24)
                    var v26 uint = v6
                    var v27 uint16 = 137
                    _ = v26
                    _ = v27
                }
            default:
                if (4294967295 > 0) {
                    var v28 int64 = -1
                    if ((!true) || v24) {
                        v2 -= v13
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        if false {
                            var v29 int8 = -1
                            _ = v29
                        }
                    }
                    _ = v28
                } else {
                    var v30 string = "\n\t"
                    for fi31 := 0; fi31 < 14; fi31++ {
                        if (!(v6 <= uint(v3))) {
                            var v32 int64 = -1
                            _ = v32
                        }
                        fmt.Println(v1, v13)
                        switch "\n\t" {
                        case "a":
                            v1 = (!v1)
                        default:
                            if len(v22) > 0 {
                                v22[0] = 1
                            }
                            switch "\n\t" {
                            case "hello":
                                var v33 int = fi23
                                _ = v33
                            }
                            var v34 int16 = -1
                            _ = v34
                        }
                        _ = fi31
                    }
                    _ = v30
                }
            }
            _ = fi23
            _ = v24
        }
        var v35 bool = v1
        _ = v35
    } else {
        _ = len(v22)
        var v36 string = "hello"
        _ = v36
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        switch "" {
        case "\n\t":
            fmt.Println(fi37, v13, fi37)
        case "hello":
            v3 += (v3 << (((^(v6))) & 63))
        case "foo bar":
            switch "" {
            case "":
                var v38 uint32 = 1
                _ = v38
            case "foo bar":
                v3 += (v3 * 183)
            case "\n\t":
                var v39 uint64 = 2
                _ = v39
            }
        default:
            fmt.Println(fi37, v13)
        }
        _ = fi37
    }
    switch 2147483647 {
    case 4294967295:
        var v40 int32 = -472
        _ = v40
    }
    fmt.Println(v6)
    fmt.Println(v6, v3)
    for fi41 := 0; fi41 < 16; fi41++ {
        v22 = append(v22, fi41)
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
