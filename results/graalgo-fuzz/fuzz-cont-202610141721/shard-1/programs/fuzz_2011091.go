package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int16; F1 string; F2 int8 }
func genFn0() bool {
    v4 := []int{}
    switch -2147483648 {
    case -9223372036854775808:
        v5 := GenS1{F0: 2147483647, F1: 42, F2: 966}
        _ = v5
    case 32767:
        for fi6 := 0; fi6 < 7; fi6++ {
            switch "foo bar" {
            case "a":
                fmt.Println(fi6)
            }
            for fi7 := 0; fi7 < 13; fi7++ {
                if false {
                    _ = len(v4)
                    _ = len(v4)
                } else {
                    var v8 int32 = -2147483648
                    fmt.Println(fi6)
                    var v9 string = "hello"
                    _ = v8
                    _ = v9
                }
                _ = fi7
            }
            _ = fi6
        }
    }
    _ = v4
    return ((-94 <= 96) || (!(18446744073709551615 < 0)))
}
func genFn1() int32 {
    var v10 uint16 = 0
    if false {
        if false {
            if true {
                v11 := GenS2{F0: -2147483648, F1: 828}
                v12 := []int{-401, 89, -506}
                fmt.Println(v10)
                _ = v11
                _ = v12
            }
            var v13 uint64 = 2
            _ = v13
        }
    }
    _ = v10
    return -1
}

func main() {
    var v1 uint32 = 0
    var v2 int8 = 1
    var v3 uint16 = 42
    switch "" {
    case "":
        v14 := []int{9223372036854775807}
        _ = v14
    }
    v15 := GenS2{F0: genFn1(), F1: 341}
    v1 -= v1
    var v16 string = ("\n\t" + "a")
    var v17 int16 = 1000
    var v18 string = (("\n\t" + v16) + "\n\t")
    v19 := GenS2{F0: -2147483648, F1: (v17 % (((v17 << ((42) & 63))) | 1))}
    v18 = v16
    switch 32767 {
    case -9223372036854775808:
        var v20 int64 = 42
        _ = v20
    case 4294967295:
        switch 4294967295 {
        case 32767:
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 uint64 = 0
                if (!false) {
                    switch ((v18 + "") + v18) {
                    case "a":
                        fmt.Println(fi21)
                    case "hello":
                        v23 := GenS2{F0: genFn1(), F1: v17}
                        _ = v23
                    case "":
                        fmt.Println(v17, v16)
                    default:
                        for fi24 := 0; fi24 < 17; fi24++ {
                            v25 := []int{(^(fi21))}
                            if ((-(v2)) > v2) {
                                var v26 int16 = (-(v17))
                                _ = v26
                            }
                            var v27 uint64 = v22
                            var v28 string = "a"
                            var v29 uint = 18446744073709551615
                            _ = fi24
                            _ = v25
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    }
                } else {
                    var v30 bool = true
                    for fi31 := 0; fi31 < 1; fi31++ {
                        v18 = "foo bar"
                        var v32 bool = (!v30)
                        var v33 int64 = int64(v17)
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    switch "" {
                    case "":
                        v34 := GenS3{F0: v17, F1: (v18 + ((v18 + v18) + ("hello" + ((v18 + (v18 + v18)) + (v18 + "\n\t"))))), F2: v2}
                        _ = v34
                    case "hello":
                        switch (fi21 ^ (fi21 + 4294967295)) {
                        case -9223372036854775808:
                            v35 := []int{-494}
                            _ = v35
                        case 42:
                            var v36 uint = 524
                            _ = v36
                        case -884:
                            var v37 int32 = 42
                            _ = v37
                        }
                    case "a":
                        var v38 uint32 = (v1 >> ((18446744073709551615) & 63))
                        _ = v38
                    }
                    _ = v30
                }
                v39 := []int{int(v3), -342, (-(fi21)), fi21}
                _ = fi21
                _ = v22
                _ = v39
            }
        }
    case 0:
        var v40 int8 = 42
        _ = v40
    default:
        v18 = (v18 + "")
    }
    v3 = v3
    fmt.Println(v1)
    switch -943 {
    case -1:
        var v41 uint64 = 445
        _ = v41
    case 42:
        var v42 int16 = 749
        _ = v42
    case 128:
        if false {
            var v43 float64 = -2.25
            _ = v43
        }
    }
    var v44 uint16 = 65535
    if (!true) {
        var v45 int16 = -281
        var v46 string = (v18 + v18)
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v44
}
