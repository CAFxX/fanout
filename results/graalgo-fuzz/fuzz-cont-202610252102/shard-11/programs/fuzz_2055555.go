package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint; F2 int32 }
func genFn0(p0 int8) int16 {
    v4 := []int{}
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 42, F1: 916}
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 string = ("" + ((("\n\t" + "hello") + "\n\t") + ((((("\n\t" + "a") + "a") + ("\n\t" + "\n\t")) + "a") + "")))
            v9 := GenS1{F0: 1, F1: 0}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(p0)
    fmt.Println(p0)
    var v10 int16 = 42
    _ = p0
    _ = v10
    _ = v4
    return (v10 - (v10 | v10))
}

func main() {
    var v1 uint32 = 42
    var v2 uint = 4294967295
    var v3 uint32 = (v1 - v1)
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: 1, F1: 65535}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            v2 -= v2
            v15 := GenS1{F0: -702, F1: 1}
            fmt.Println(v1, fi13, fi14)
            _ = fi14
            _ = v15
        }
        switch fi13 {
        case 127:
            v16 := GenS2{F0: -252, F1: 607, F2: v2}
            _ = v16
        case -9223372036854775808:
            var v17 float32 = 3.14159
            _ = v17
        default:
            v3 -= uint32(v2)
        }
        if true {
            v18 := GenS3{F0: 1}
            v19 := GenS2{F0: 496, F1: uint(v3), F2: (^(v2))}
            _ = v18
            _ = v19
        }
        _ = fi13
    }
    if true {
        v1 -= (v3 % (((v1 & (v3 & v3))) | 1))
        for fi20 := 0; fi20 < 18; fi20++ {
            v2 = v2
            _ = fi20
        }
        if false {
            v21 := []int{42}
            _ = v21
        } else {
            var v22 uint8 = 0
            var v23 int = 127
            _ = v22
            _ = v23
        }
    }
    if (0 == 42) {
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 int16 = -32768
            v2 += uint(v3)
            _ = fi24
            _ = v25
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := []int{127, (fi26 + (-(fi26))), fi26, fi26}
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            _ = fi26
            _ = v27
        }
    }
    v1 = (v3 - v1)
    v2 -= v2
    var v28 int16 = 0
    switch 32768 {
    case 1:
        v3 -= (v1 / ((42) | 1))
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := []int{(fi29 >> (((v2 << (((v2 << ((v2) & 63))) & 63))) & 63)), 4294967295}
        v3 += 1
        switch 809 {
        case -9223372036854775808:
            var v31 uint16 = 32767
            _ = v31
        default:
            var v32 uint = v2
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        v3 += 2147483648
        _ = fi33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        var v35 uint64 = 345
        if ((3.14159 / 1.0) > 3.14159) {
            if (!(v1 > v1)) {
                var v36 string = "\n\t"
                fmt.Println(v2)
                _ = v36
            }
            switch ("\n\t" + "\n\t") {
            case "a":
                for fi37 := 0; fi37 < 12; fi37++ {
                    v28 = int16(fi34)
                    var v38 uint64 = (^(v35))
                    v35 = (v38 ^ v38)
                    fmt.Println(v35)
                    v39 := []int{}
                    v35 = v38
                    fmt.Println(v1, v2)
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            case "\n\t":
                v40 := GenS1{F0: 42, F1: 65535}
                _ = v40
            case "foo bar":
                fmt.Println(fi34, v3, v3)
            default:
                v41 := GenS1{F0: int64(v35), F1: 0}
                _ = v41
            }
        } else {
            var v42 uint32 = (^(v1))
            fmt.Println(v1, v28, fi34)
            for fi43 := 0; fi43 < 10; fi43++ {
                if false {
                    var v44 float32 = 1e10
                    if ((32767 >= 1) && (0.0 > 0.0)) {
                        var v45 uint16 = 567
                        if false {
                            v1 -= (v3 + v42)
                        }
                        v28 += (v28 - (v28 & v28))
                        _ = v45
                    }
                    v46 := GenS3{F0: 1}
                    _ = v44
                    _ = v46
                }
                if (!(-1 > -5)) {
                    switch (fi34 << ((4294967295) & 63)) {
                    case 2147483647:
                        fmt.Println(v42, v1)
                        v35 -= 2
                        v1 += v1
                    }
                    v47 := GenS4{F0: (v1 * 2), F1: 18446744073709551615, F2: 1}
                    _ = v47
                } else {
                    fmt.Println(fi34, v42, fi34)
                }
                v48 := GenS4{F0: v3, F1: (v2 >> (((^(v2))) & 63)), F2: -2147483648}
                _ = fi43
                _ = v48
            }
            _ = v42
        }
        for fi49 := 0; fi49 < 14; fi49++ {
            v35 -= (^(v35))
            _ = fi49
        }
        _ = fi34
        _ = v35
    }
    v50 := []int{-2147483648}
    switch -899 {
    case -936:
        var v51 uint8 = 128
        _ = v51
    case 0:
        var v52 int = 128
        _ = v52
    default:
        var v53 uint = (^(v2))
        _ = v53
    }
    for fi54 := 0; fi54 < 15; fi54++ {
        v55 := GenS4{F0: 2, F1: v2, F2: 42}
        var v56 uint64 = 4294967295
        v1 = uint32(v56)
        _ = fi54
        _ = v55
        _ = v56
    }
    if (false || false) {
        v57 := GenS4{F0: 720, F1: v2, F2: -954}
        var v58 int64 = -168
        _ = v57
        _ = v58
    }
    switch "a" {
    case "foo bar":
        for fi59 := 0; fi59 < 15; fi59++ {
            v2 -= 1
            v2 = 65535
            _ = fi59
        }
    case "":
        fmt.Println(v28, v28, v28)
    }
    var v60 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v50
    _ = v60
}
