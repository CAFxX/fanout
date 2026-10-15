package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int16 {
    var v4 uint32 = 2
    switch 0 {
    case -9223372036854775808:
        switch -9223372036854775808 {
        case 9223372036854775807:
            if false {
                switch -741 {
                case 2147483647:
                    v4 -= (^(v4))
                case 9223372036854775807:
                    var v5 uint32 = 1
                    _ = v5
                case -318:
                    var v6 uint32 = 65535
                    _ = v6
                }
                var v7 bool = ((42 != 2) && true)
                var v8 bool = (28 <= 42)
                _ = v7
                _ = v8
            } else {
                var v9 float64 = (0.0 - 0.0)
                _ = v9
            }
        }
    case 1:
        v10 := []int{-1, 532, 106}
        _ = v10
    case -2147483648:
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint8 = 64
            _ = fi11
            _ = v12
        }
    }
    _ = v4
    return 42
}
func genFn1(p0 int32, p1 bool, p2 int8) uint64 {
    if p1 {
        var v13 uint32 = 187
        _ = v13
    } else {
        fmt.Println(p1)
        for fi14 := 0; fi14 < 17; fi14++ {
            switch (-(fi14)) {
            case 9223372036854775807:
                var v15 bool = false
                _ = v15
            case -1:
                var v16 int32 = p0
                _ = v16
            }
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 bool = false
                var v19 int16 = 0
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi14
        }
    }
    p1 = (2147483647 <= 2)
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint16 = 65535
    var v2 int8 = 1
    var v3 string = (((((((((("" + "a") + "foo bar") + (("hello" + ("" + "hello")) + "foo bar")) + "foo bar") + ("a" + "\n\t")) + "") + "foo bar") + (("" + "foo bar") + "\n\t")) + "a") + "\n\t")
    var v20 int = 4294967295
    var v21 int32 = 0
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS1{F0: 1}
        _ = fi22
        _ = v23
    }
    v20 = 700
    v20 = 9223372036854775807
    var v24 uint8 = 255
    var v25 int16 = 42
    if (!true) {
        v25 = -282
        if (true || true) {
            var v26 int = v20
            v27 := []int{-816, v20, (v20 + v26), v20}
            if false {
                if false {
                    switch v3 {
                    case "\n\t":
                        for fi28 := 0; fi28 < 15; fi28++ {
                            v29 := GenS1{F0: int64(v2)}
                            _ = fi28
                            _ = v29
                        }
                    case "":
                        v21 -= (v21 >> ((18446744073709551615) & 63))
                    case "a":
                        v1 += (v1 - 617)
                    default:
                        if len(v27) > 0 {
                            v27[0] = v20
                        }
                    }
                    for fi30 := 0; fi30 < 12; fi30++ {
                        switch 128 {
                        case -757:
                            for fi31 := 0; fi31 < 8; fi31++ {
                                if false {
                                    var v32 string = v3
                                    _ = v32
                                } else {
                                    var v33 uint8 = v24
                                    _ = v33
                                }
                                if len(v27) > 0 {
                                    v27[0] = v26
                                }
                                _ = fi31
                            }
                        case -946:
                            var v34 int64 = -614
                            _ = v34
                        }
                        var v35 int16 = 1000
                        _ = fi30
                        _ = v35
                    }
                    var v36 bool = true
                    _ = v36
                } else {
                    var v37 uint16 = uint16(v24)
                    _ = v37
                }
            } else {
                for fi38 := 0; fi38 < 9; fi38++ {
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                    var v39 float64 = -2.25
                    _ = fi38
                    _ = v39
                }
                var v40 int8 = 42
                if len(v27) > 0 {
                    v27[0] = (v26 / ((v26) | 1))
                }
                _ = v40
            }
            _ = v26
            _ = v27
        }
    }
    var v41 float64 = ((0.1 - ((((3.14159 - (0.0 * 1.5)) * 0.1) * ((1.5 / 1.5) / 3.14159)) / (0.0 - 1.5))) + 1e10)
    var v42 int = v20
    for fi43 := 0; fi43 < 8; fi43++ {
        var v44 float64 = (float64((0.0 - 0.0)) - v41)
        var v45 float64 = float64(3.14159)
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v41
    _ = v42
}
