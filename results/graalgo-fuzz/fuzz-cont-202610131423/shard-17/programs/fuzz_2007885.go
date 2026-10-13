package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int32, p1 uint32, p2 int8) uint32 {
    p0 = -863
    v7 := GenS1{F0: 358, F1: -1, F2: p0}
    var v8 uint64 = 2
    if (true && true) {
        p2 = (p2 << ((16) & 63))
        var v9 uint = 535
        var v10 string = "foo bar"
        _ = v10
        _ = v9
    } else {
        var v11 uint32 = p1
        fmt.Println(v11, p0, p1)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 / ((p1) | 1))
}
func genFn1() uint64 {
    if (true || (((!false) || true) && (42 == 20))) {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := []int{(fi12 + (fi12 + 4294967295))}
            _ = fi12
            _ = v13
        }
    } else {
        var v14 bool = ("a" > "a")
        var v15 uint = 0
        _ = v14
        _ = v15
    }
    var v16 uint32 = 65535
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := []int{(fi17 + fi17)}
        var v19 uint16 = 2
        if (67 < 127) {
            v20 := GenS2{F0: 686}
            if (false || (-26 < fi17)) {
                for fi21 := 0; fi21 < 18; fi21++ {
                    var v22 int = 2147483647
                    v23 := GenS3{F0: -2147483648}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                var v24 int8 = -128
                _ = v24
            }
            _ = v20
        } else {
            switch 42 {
            case 2147483647:
                var v25 int16 = 1000
                _ = v25
            case -2147483648:
                for fi26 := 0; fi26 < 18; fi26++ {
                    fmt.Println(fi17, fi26)
                    if (863 <= 518) {
                        var v27 int = 0
                        _ = v27
                    } else {
                        if len(v18) > 0 {
                            v18[0] = -9223372036854775808
                        }
                        v19 = (v19 | v19)
                    }
                    _ = fi26
                }
            }
            for fi28 := 0; fi28 < 7; fi28++ {
                switch 42 {
                case -192:
                    var v29 int16 = 1000
                    _ = v29
                case -2147483648:
                    v30 := GenS2{F0: 1}
                    _ = v30
                case -9223372036854775808:
                    v16 = v16
                default:
                    if (false || (18446744073709551615 < 0)) {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    }
                }
                v19 = 255
                if false {
                    var v31 string = "foo bar"
                    _ = v31
                }
                if (25 != -1) {
                    var v32 uint32 = (^(v16))
                    var v33 uint64 = 2
                    v19 = 2
                    var v34 float32 = 1e10
                    _ = v32
                    _ = v33
                    _ = v34
                } else {
                    var v35 uint8 = uint8(v16)
                    if len(v18) > 0 {
                        v18[0] = (fi28 / ((fi28) | 1))
                    }
                    v16 -= v16
                    _ = v35
                }
                _ = fi28
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                switch (fi17 & (fi17 * (fi17 - (fi36 / ((2147483647) | 1))))) {
                case 128:
                    v19 = v19
                case 291:
                    fmt.Println(fi36)
                }
                if false {
                    var v37 int64 = 1
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    var v38 float64 = float64(1e10)
                    _ = v37
                    _ = v38
                }
                v18 = append(v18, fi17)
                var v39 float64 = ((3.14159 + (0.1 + 0.1)) / 1.0)
                _ = fi36
                _ = v39
            }
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = v16
    return 0
}

func main() {
    var v1 int16 = -32768
    v2 := []string{"a", ("\n\t" + "a"), ("foo bar" + "\n\t"), "\n\t"}
    var v3 uint8 = 1
    var v4 uint = 4294967295
    var v5 float64 = 0.0
    var v6 uint16 = 255
    var v40 uint64 = genFn1()
    if true {
        fmt.Println(v40, v5)
    }
    if (((-1 == -9223372036854775808) || true) || (!true)) {
        for fi41 := 0; fi41 < 3; fi41++ {
            switch (fi41 % ((fi41) | 1)) {
            case -222:
                v2 = append(v2, ("foo bar" + ("" + "hello")))
            case 128:
                v1 -= (v1 | v1)
            case -376:
                var v42 int8 = 100
                _ = v42
            default:
                var v43 uint64 = v40
                _ = v43
            }
            _ = fi41
        }
    }
    v3 = 6
    v1 = (v1 - v1)
    switch 4294967295 {
    case 32767:
        v40 = (^(v40))
    case 598:
        var v44 uint16 = v6
        _ = v44
    case 2147483647:
        if ((v5 + v5) < v5) {
            fmt.Println(v4, v1, v40)
            for fi45 := 0; fi45 < 18; fi45++ {
                v46 := GenS3{F0: -2147483648}
                _ = fi45
                _ = v46
            }
        }
    }
    v2 = append(v2, "")
    switch 128 {
    case -583:
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    case -48:
        var v47 uint64 = uint64(v6)
        _ = v47
    case 32767:
        v48 := GenS3{F0: 627}
        _ = v48
    default:
        if ((true && ((!(0 == -1)) || (3.14159 == 1e10))) && false) {
            v1 = v1
        } else {
            fmt.Println(v1)
        }
    }
    if false {
        fmt.Println(v4)
    } else {
        v1 += (^(v1))
    }
    var v49 float64 = v5
    fmt.Println(v1)
    if len(v2) > 0 {
        v2[0] = "a"
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v49
    _ = v5
    _ = v6
}
