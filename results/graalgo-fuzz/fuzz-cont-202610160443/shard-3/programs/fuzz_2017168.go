package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64, p1 int8, p2 uint64) uint16 {
    p2 = 323
    if (!(!(!(!(!true))))) {
        p1 = p1
    }
    v7 := []int{2147483647, -369, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 2147483647
    v3 := []uint32{(^(v2)), (v2 - 42), (v2 >> ((65535) & 63)), v2}
    var v4 int32 = -785
    var v5 uint32 = v2
    v6 := []uint32{65535, v2}
    v1 += v1
    for fi8 := 0; fi8 < 8; fi8++ {
        if (!(65535 != 850)) {
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 uint32 = v5
                var v11 string = "foo bar"
                var v12 int16 = -1
                var v13 int64 = 42
                v14 := GenS1{F0: 1}
                fmt.Println(v2, v1)
                fmt.Println(v5)
                var v15 uint32 = 0
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 float64 = ((float64(-2.25) / 0.1) * (float64(3.14159) * 1e10))
            if true {
                var v17 float64 = (((0.0 - float64(float32((1.5 + 1.5)))) - (v16 / 1.5)) / v16)
                var v18 uint64 = 656
                switch (fi8 & (fi8 << ((1) & 63))) {
                case 669:
                    var v19 float64 = v17
                    _ = v19
                }
                _ = v17
                _ = v18
            } else {
                var v20 float32 = 1e10
                _ = v20
            }
            _ = v16
        }
        _ = fi8
    }
    switch 354 {
    case 660:
        for fi21 := 0; fi21 < 14; fi21++ {
            for fi22 := 0; fi22 < 19; fi22++ {
                v23 := GenS1{F0: 455}
                var v24 bool = true
                switch ("" + ("\n\t" + "foo bar")) {
                case "\n\t":
                    v5 = (v2 + v2)
                case "a":
                    var v25 int64 = -991
                    _ = v25
                }
                v4 = v4
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = len(v3)
            fmt.Println(v4, v2)
            _ = fi21
        }
    }
    if (!false) {
        _ = len(v6)
    } else {
        v26 := GenS1{F0: 65535}
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        fmt.Println(v1)
        var v28 int32 = v4
        _ = fi27
        _ = v28
    }
    var v29 int64 = 42
    fmt.Println(v29)
    if len(v3) > 0 {
        v3[0] = (v5 | (v2 * (^(v2))))
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        switch (fi30 << ((1) & 63)) {
        case 1:
            for fi31 := 0; fi31 < 3; fi31++ {
                v29 = 418
                _ = fi31
            }
        case 32767:
            v32 := GenS1{F0: 4294967295}
            _ = v32
        }
        if (-2.25 > float64((3.14159 / 1.5))) {
            fmt.Println(v2, v4)
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi30
    }
    v33 := GenS1{F0: 85}
    for fi34 := 0; fi34 < 10; fi34++ {
        switch (-(fi34)) {
        case -279:
            if (true || (!((9223372036854775807 >= (v29 | v29)) && (835 != 42)))) {
                if false {
                    v2 = (v2 * (v5 - (v2 >> ((18446744073709551615) & 63))))
                    v29 = v29
                    var v35 uint32 = (v2 + v2)
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 12; fi36++ {
                        v4 += 42
                        _ = fi36
                    }
                }
                v6 = append(v6, v2)
            }
        }
        fmt.Println(v2, fi34)
        _ = fi34
    }
    v37 := GenS1{F0: uint(v5)}
    switch "" {
    case "":
        var v38 bool = (v4 >= v4)
        _ = v38
    case "foo bar":
        fmt.Println(v1, v1)
    case "a":
        for fi39 := 0; fi39 < 8; fi39++ {
            for fi40 := 0; fi40 < 8; fi40++ {
                var v41 uint64 = 1
                _ = fi40
                _ = v41
            }
            for fi42 := 0; fi42 < 14; fi42++ {
                var v43 bool = true
                v43 = false
                _ = fi42
                _ = v43
            }
            fmt.Println(v2, fi39)
            _ = fi39
        }
    }
    var v44 int8 = v1
    for fi45 := 0; fi45 < 12; fi45++ {
        var v46 uint16 = 32767
        _ = fi45
        _ = v46
    }
    v1 = (v1 + (v44 % ((v44) | 1)))
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
