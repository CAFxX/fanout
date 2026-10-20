package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint8 }
func genFn0() int32 {
    var v6 bool = false
    v7 := []int{-9223372036854775808, -144}
    var v8 int16 = -32768
    _ = v6
    _ = v7
    _ = v8
    return -769
}
func genFn1(p0 int16, p1 uint32, p2 int32) string {
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := []int{fi9, fi9, (fi9 - fi9), (fi9 << ((65535) & 63))}
        p0 = (-(p0))
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        p0 -= -61
        _ = fi11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := []int{fi12, fi12, fi12, (fi12 + fi12)}
        var v14 uint = 4294967295
        _ = fi12
        _ = v13
        _ = v14
    }
    if false {
        var v15 int32 = p2
        v16 := []int{-314, -891, -800, 42}
        v15 = v15
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]bool)
    var v3 int64 = -66
    var v4 int32 = 1
    var v5 uint = 47
    switch 42 {
    case 42:
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(v4, v3, v1)
            _ = fi17
        }
    case -9223372036854775808:
        v18 := []int{9223372036854775807, -289, 32767}
        _ = v18
    }
    switch (("\n\t" + "") + "hello") {
    case "foo bar":
        v19 := []int{}
        _ = v19
    }
    if false {
        v4 += genFn0()
        switch -900 {
        case 2147483647:
            v20 := []int{}
            _ = v20
        }
    }
    switch -235 {
    case 2147483647:
        v21 := GenS1{F0: true, F1: genFn0(), F2: 216}
        _ = v21
    }
    fmt.Println(v5, v3)
    switch -887 {
    case -1:
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 uint64 = 0
            v4 -= (v4 / ((v4) | 1))
            _ = fi22
            _ = v23
        }
    case 0:
        for fi24 := 0; fi24 < 13; fi24++ {
            if false {
                var v25 float32 = -2.25
                _ = v25
            }
            _ = fi24
        }
    case 99:
        v5 -= (v5 << ((v5) & 63))
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        switch (fi26 | (fi26 << (((v5 % (((v5 / (((^(v5))) | 1))) | 1))) & 63))) {
        case 9223372036854775807:
            var v27 uint8 = 0
            _ = v27
        case -933:
            v4 -= (-(v4))
        }
        if ("" != "hello") {
            switch (fi26 << (((v5 & 4294967295)) & 63)) {
            case 414:
                switch fi26 {
                case 42:
                    for fi28 := 0; fi28 < 14; fi28++ {
                        v29 := GenS1{F0: true, F1: 0, F2: 255}
                        var v30 float64 = -2.25
                        fmt.Println(fi28)
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case 0:
                    var v31 string = "\n\t"
                    _ = v31
                }
            }
            var v32 int64 = 9223372036854775807
            _ = v32
        } else {
            if (!true) {
                for fi33 := 0; fi33 < 16; fi33++ {
                    if true {
                        var v34 uint = (v5 ^ v5)
                        var v35 int16 = int16(v3)
                        v36 := []int{}
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        v37 := GenS1{F0: true, F1: int32(v3), F2: 1}
                        _ = v37
                    }
                    var v38 int = fi33
                    v39 := GenS1{F0: (v4 == -1), F1: (v4 % ((v4) | 1)), F2: 128}
                    _ = fi33
                    _ = v38
                    _ = v39
                }
            }
            var v40 uint8 = 42
            v4 -= (v4 / ((genFn0()) | 1))
            _ = v40
        }
        _ = fi26
    }
    fmt.Println(v5, v3)
    if (18446744073709551615 < v5) {
        var v41 string = "foo bar"
        fmt.Println(v41, v3, v41)
        v42 := []int{42, 127, 2147483647, 32768}
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
