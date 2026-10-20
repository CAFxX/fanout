package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    var v1 int = 1
    var v2 uint64 = 0
    var v3 int64 = 478
    var v4 uint = 255
    var v5 int64 = (v3 + v3)
    v6 := GenS1{F0: (("hello" + "\n\t") + ("" + "a"))}
    var v7 int16 = 32767
    for fi8 := 0; fi8 < 17; fi8++ {
        switch (fi8 + 32767) {
        case -878:
            var v9 bool = ((v7 < v7) || (-128 > -1))
            _ = v9
        case -2147483648:
            v10 := []int{(fi8 ^ fi8)}
            _ = v10
        }
        switch fi8 {
        case 0:
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := []int{42, (^(v1))}
                switch (fi8 + (v1 * 146)) {
                case -419:
                    v12 = append(v12, fi11)
                case 1:
                    v13 := GenS1{F0: "hello"}
                    _ = v13
                }
                var v14 uint8 = 127
                _ = fi11
                _ = v12
                _ = v14
            }
        default:
            var v15 string = "foo bar"
            _ = v15
        }
        v16 := []int{v1, (v1 << ((v4) & 63)), v1}
        _ = fi8
        _ = v16
    }
    v7 += v7
    switch "a" {
    case "hello":
        var v17 int16 = -270
        _ = v17
    case "\n\t":
        v18 := GenS1{F0: "hello"}
        _ = v18
    default:
        var v19 uint32 = 42
        _ = v19
    }
    var v20 uint16 = 65535
    var v21 int32 = 42
    v2 += 182
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := GenS1{F0: ("foo bar" + "\n\t")}
        if ((!(1e10 == 0.1)) && false) {
            var v24 int8 = 54
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        v20 = v20
        if (2147483648 != uint32(v4)) {
            var v26 float32 = ((0.0 / 3.14159) / float32(((0.1 - float64(1e10)) / 3.14159)))
            var v27 int = v1
            _ = v26
            _ = v27
        }
        _ = fi25
    }
    v21 = int32(v20)
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := []int{}
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v4 = 0
        _ = fi30
    }
    switch v1 {
    case 961:
        var v31 uint16 = (v20 ^ 32767)
        _ = v31
    case 32768:
        if ((2 != 0) && true) {
            var v32 int16 = v7
            for fi33 := 0; fi33 < 2; fi33++ {
                fmt.Println(v7, v7, v5)
                for fi34 := 0; fi34 < 1; fi34++ {
                    v35 := GenS1{F0: "foo bar"}
                    _ = fi34
                    _ = v35
                }
                for fi36 := 0; fi36 < 12; fi36++ {
                    v37 := []int{(-(v1)), -9223372036854775808}
                    fmt.Println(fi36)
                    _ = fi36
                    _ = v37
                }
                _ = fi33
            }
            _ = v32
        } else {
            var v38 uint16 = uint16(v2)
            var v39 int8 = 0
            _ = v38
            _ = v39
        }
    case 499:
        var v40 int64 = 0
        _ = v40
    }
    switch v1 {
    case 475:
        switch "" {
        case "":
            var v41 uint64 = v2
            _ = v41
        }
    default:
        switch "\n\t" {
        case "foo bar":
            v42 := GenS1{F0: ("a" + "")}
            _ = v42
        }
    }
    fmt.Println(v2, v4)
    switch "\n\t" {
    case "":
        if (((v2 & v2) < 474) || ((v20 + v20) < v20)) {
            v3 += -574
            var v43 bool = ((v4 - 18446744073709551615) >= 255)
            _ = v43
        }
    case "hello":
        var v44 uint = v4
        _ = v44
    default:
        v45 := GenS2{F0: (v3 | (v3 >> ((496) & 63))), F1: int32(v5)}
        _ = v45
    }
    v46 := []int{4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v7
}
