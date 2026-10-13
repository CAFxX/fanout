package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 string; F1 int; F2 uint }
type GenS3 struct { F0 int8; F1 uint8; F2 uint64 }

func main() {
    var v1 string = ("hello" + (("a" + "\n\t") + "\n\t"))
    var v2 float32 = (1.5 - 3.14159)
    var v3 int16 = 1
    var v4 bool = (!(!false))
    var v5 int64 = 1
    if (!((!(false || (!v4))) || (2147483647 >= 965))) {
        switch "hello" {
        case "":
            var v6 bool = false
            _ = v6
        default:
            v3 = -625
        }
        var v7 float32 = v2
        if (42 > 2147483647) {
            var v8 uint = 550
            _ = v8
        } else {
            v7 = v2
            if v4 {
                var v9 float64 = (0.1 + (1e10 * (-2.25 - 3.14159)))
                _ = v9
            }
            var v10 int16 = v3
            _ = v10
        }
        _ = v7
    }
    if ((60 >= 1) || ((!(1 < v5)) && (!(-1 >= -812)))) {
        v11 := GenS1{F0: (-(v3)), F1: 42, F2: 255}
        v12 := GenS1{F0: 0, F1: 42, F2: 35}
        v13 := []int{}
        _ = v11
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: v3, F1: 128, F2: 95}
        for fi15 := 0; fi15 < 1; fi15++ {
            if (0 < 543) {
                for fi16 := 0; fi16 < 10; fi16++ {
                    fmt.Println(fi16, v3, v5)
                    _ = fi16
                }
                v17 := GenS1{F0: (v3 >> ((4294967295) & 63)), F1: (-(fi15)), F2: 2}
                v18 := []int{(fi15 | (fi15 << ((2) & 63))), fi15}
                _ = v17
                _ = v18
            }
            if v4 {
                v2 = v2
            }
            if v4 {
                v1 = (v1 + v1)
            }
            _ = fi15
        }
        if v4 {
            v19 := GenS2{F0: ("hello" + v1), F1: -755, F2: 255}
            v2 = v2
            _ = v19
        } else {
            fmt.Println(v4, v5)
            v3 = 1000
        }
        _ = v14
    }
    fmt.Println(v2)
    if (301 == 699) {
        fmt.Println(v3, v4)
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        v2 -= v2
        var v21 int = fi20
        var v22 int = v21
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(v1, v3, v3)
    if (!(!true)) {
        v23 := []int{32767}
        switch -2147483648 {
        case 128:
            var v24 int = 32768
            _ = v24
        case -9223372036854775808:
            var v25 bool = ((v4 && v4) && true)
            _ = v25
        }
        v26 := GenS2{F0: (v1 + (((("\n\t" + (v1 + (v1 + "hello"))) + ((v1 + ("a" + "a")) + "a")) + ("\n\t" + v1)) + "")), F1: 371, F2: 199}
        _ = v23
        _ = v26
    }
    if (100 < 39) {
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := []int{fi27, fi27, (fi27 * fi27), fi27}
            _ = len(v28)
            _ = fi27
            _ = v28
        }
        v1 = v1
    }
    var v29 int8 = -128
    if (65535 != 255) {
        var v30 uint32 = 4294967295
        var v31 float64 = float64(float32(1e10))
        var v32 uint64 = uint64(v30)
        _ = v30
        _ = v31
        _ = v32
    } else {
        if v4 {
            for fi33 := 0; fi33 < 15; fi33++ {
                v34 := []int{fi33, fi33}
                v35 := GenS3{F0: v29, F1: 2, F2: 1}
                var v36 float64 = 1.5
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            v5 += v5
            _ = fi37
        }
        v38 := GenS1{F0: int16(v29), F1: 42, F2: 128}
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        v2 = 1e10
        var v40 int16 = v3
        v5 = v5
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        v1 = v1
        _ = fi41
    }
    fmt.Println(v4)
    if ((2147483647 != 2) || v4) {
        v42 := []int{}
        _ = v42
    } else {
        v2 += (0.0 + float32(-2.25))
        v43 := []int{}
        v1 = v1
        _ = v43
    }
    var v44 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
