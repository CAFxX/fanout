package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }
func genFn0() int16 {
    if false {
        v5 := GenS1{F0: 2, F1: 159, F2: -9223372036854775808}
        _ = v5
    }
    var v6 bool = false
    for fi7 := 0; fi7 < 4; fi7++ {
        v6 = (42 == -128)
        _ = fi7
    }
    if v6 {
        v8 := []int{}
        _ = len(v8)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := []int{(^(fi9))}
            var v11 uint64 = 622
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := []int{-896, 1, 327}
        if ((!v6) && v6) {
            v6 = v6
        } else {
            switch 4294967295 {
            case 9223372036854775807:
                v13 := GenS1{F0: 1, F1: 2, F2: -9223372036854775808}
                _ = v13
            }
            switch "" {
            case "\n\t":
                v6 = false
            case "hello":
                for fi14 := 0; fi14 < 17; fi14++ {
                    v6 = v6
                    _ = fi14
                }
            }
        }
        _ = v12
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 int32 = 0
        v16 += -2147483648
        _ = fi15
        _ = v16
    }
    _ = v6
    return -32768
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = 641
    var v3 int8 = -24
    var v4 int64 = -731
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int = (fi17 % ((32767) | 1))
        v2 -= v4
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 uint8 = 42
            if (true || false) {
                v4 = int64(fi17)
                fmt.Println(v20, v3, fi19)
                v21 := []int{(v18 * (fi19 ^ (fi17 / (((fi19 ^ fi19)) | 1)))), (-(fi17))}
                var v22 bool = true
                _ = v21
                _ = v22
            } else {
                switch (v18 | 32768) {
                case -648:
                    switch (fi17 >> ((326) & 63)) {
                    case 9223372036854775807:
                        fmt.Println(v4, v20)
                    case 979:
                        if true {
                            var v23 uint8 = v20
                            _ = v23
                        }
                    case -2147483648:
                        v24 := GenS2{F0: v3, F1: 127}
                        var v25 int32 = 544
                        switch ("" + ("" + "")) {
                        case "foo bar":
                            v26 := []int{(^(fi17)), 4294967295}
                            v27 := GenS2{F0: (v3 >> ((259) & 63)), F1: v18}
                            v1 += (^(v1))
                            _ = v26
                            _ = v27
                        case "\n\t":
                            var v28 bool = true
                            _ = v28
                        case "hello":
                            if (false || false) {
                                v29 := GenS1{F0: 42, F1: 42, F2: v4}
                                fmt.Println(v25, v18, v1)
                                v30 := GenS3{F0: ((v4 ^ int64(v25)) < v4), F1: (fi19 - fi17)}
                                _ = v29
                                _ = v30
                            }
                        default:
                            v3 += (v3 >> ((4294967295) & 63))
                        }
                        _ = v24
                        _ = v25
                    default:
                        v31 := GenS4{F0: 760}
                        _ = v31
                    }
                case 127:
                    var v32 bool = (544 < 2)
                    _ = v32
                case 128:
                    v33 := []int{128, 2147483647, fi17}
                    _ = v33
                }
                var v34 float32 = ((-2.25 - 1e10) + (float32(3.14159) / 1.0))
                var v35 int32 = 42
                _ = v34
                _ = v35
            }
            var v36 int16 = 1000
            if (744 < uint(v1)) {
                fmt.Println(v36, fi19)
                v36 += (v36 % ((32767) | 1))
                fmt.Println(v18, v18)
                v18 = fi19
            } else {
                fmt.Println(v4, v36)
                var v37 uint8 = (^(v20))
                v38 := GenS5{F0: (("foo bar" + "hello") + (("hello" + "\n\t") + "\n\t"))}
                _ = v37
                _ = v38
            }
            _ = fi19
            _ = v20
            _ = v36
        }
        _ = fi17
        _ = v18
    }
    v39 := GenS6{F0: 853}
    v4 += v2
    var v40 int32 = 800
    v40 = 944
    if ((v3 - 100) >= v3) {
        if (true || (127 > 1)) {
            var v41 int8 = (v3 << ((415) & 63))
            v4 = -1
            _ = v41
        } else {
            v42 := []int{}
            _ = v42
        }
    } else {
        v43 := GenS1{F0: 195, F1: 255, F2: v2}
        v3 -= v3
        _ = v43
    }
    var v44 int16 = -623
    switch 9223372036854775807 {
    case 1:
        v1 = 65535
    case 32768:
        var v45 int16 = v44
        _ = v45
    case 9223372036854775807:
        for fi46 := 0; fi46 < 1; fi46++ {
            v47 := []int{}
            v3 += v3
            v3 = v3
            _ = fi46
            _ = v47
        }
    }
    if (!true) {
        if (!((^(v40)) != v40)) {
            var v48 string = ""
            var v49 string = (v48 + (v48 + (v48 + (v48 + (v48 + v48)))))
            _ = v48
            _ = v49
        } else {
            for fi50 := 0; fi50 < 2; fi50++ {
                fmt.Println(v40, fi50, v44)
                v3 = (v3 << ((0) & 63))
                _ = fi50
            }
            for fi51 := 0; fi51 < 6; fi51++ {
                for fi52 := 0; fi52 < 5; fi52++ {
                    v53 := GenS5{F0: "a"}
                    v54 := GenS6{F0: 0}
                    var v55 float64 = ((1e10 - (1.5 / 1.5)) * 0.0)
                    _ = fi52
                    _ = v53
                    _ = v54
                    _ = v55
                }
                v3 = (v3 >> ((2) & 63))
                switch "a" {
                case "a":
                    for fi56 := 0; fi56 < 8; fi56++ {
                        fmt.Println(fi51)
                        _ = fi56
                    }
                default:
                    var v57 uint = 392
                    _ = v57
                }
                _ = fi51
            }
            if (v44 < 0) {
                fmt.Println(v44, v4)
            }
        }
    }
    switch 32767 {
    case -9223372036854775808:
        if false {
            var v58 string = "foo bar"
            switch v58 {
            case "foo bar":
                v4 = -9223372036854775808
            case "a":
                v59 := []int{4294967295, 32767, 666}
                _ = v59
            case "\n\t":
                for fi60 := 0; fi60 < 9; fi60++ {
                    if false {
                        var v61 uint64 = 828
                        _ = v61
                    }
                    _ = fi60
                }
            }
            if (9 <= 255) {
                switch 1 {
                case 116:
                    var v62 int32 = v40
                    _ = v62
                case -1:
                    var v63 bool = false
                    _ = v63
                default:
                    v64 := GenS5{F0: "foo bar"}
                    _ = v64
                }
                v65 := GenS5{F0: v58}
                v66 := []int{}
                _ = v65
                _ = v66
            }
            _ = v58
        }
    }
    v67 := GenS4{F0: v40}
    if false {
        v4 += (v2 >> ((18446744073709551615) & 63))
        v3 = (v3 ^ v3)
        switch "" {
        case "\n\t":
            var v68 uint16 = 0
            _ = v68
        case "hello":
            var v69 uint = 255
            _ = v69
        }
    }
    v4 -= int64(v1)
    fmt.Println(v44, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v67
}
