package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint64, p1 bool) int32 {
    var v7 uint8 = 120
    v8 := []int{186, 32767}
    p0 = 2
    p1 = true
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 783
}

func main() {
    var v1 bool = (42 != -32768)
    v2 := []string{}
    v3 := make(map[string]int)
    var v4 int8 = 0
    var v5 uint = 577
    v6 := make(map[string]string)
    for fi9 := 0; fi9 < 7; fi9++ {
        v4 = (v4 & (-(v4)))
        _ = fi9
    }
    v4 -= (v4 + v4)
    var v10 uint32 = 1
    for fi11 := 0; fi11 < 7; fi11++ {
        v1 = v1
        v10 = v10
        _ = fi11
    }
    switch ("\n\t" + ("a" + "a")) {
    case "foo bar":
        switch 4294967295 {
        case 127:
            switch "foo bar" {
            case "hello":
                var v12 uint8 = 85
                _ = v12
            case "foo bar":
                for fi13 := 0; fi13 < 8; fi13++ {
                    fmt.Println(v10, v5, v4)
                    delete(v6, "foo bar")
                    _ = fi13
                }
            case "":
                var v14 uint32 = v10
                _ = v14
            }
        case -1:
            if (1 <= 1) {
                fmt.Println(v10, v5)
                _ = len(v2)
                var v15 float32 = 1e10
                _ = v15
            }
        }
    }
    var v16 uint32 = v10
    var v17 uint16 = 2
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := GenS1{F0: -2147483648}
            fmt.Println(v10, fi19)
            var v21 uint64 = 2
            var v22 uint64 = (v21 % (((v21 * (v21 + v21))) | 1))
            v21 = 1
            v23 := GenS1{F0: 1}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
            _ = v23
        }
        switch fi18 {
        case 381:
            if v1 {
                for fi24 := 0; fi24 < 20; fi24++ {
                    var v25 int64 = 9223372036854775807
                    switch (fi18 & 127) {
                    case 229:
                        _ = len(v3)
                    default:
                        v25 = -1
                    }
                    _ = fi24
                    _ = v25
                }
                v10 = uint32(v5)
            } else {
                var v26 float64 = ((float64(1e10) + -2.25) / 3.14159)
                var v27 int8 = (v4 + (v4 >> (((^(v5))) & 63)))
                delete(v6, "hello")
                _ = v26
                _ = v27
            }
        case 402:
            var v28 float64 = -2.25
            _ = v28
        case 749:
            v29 := GenS1{F0: -870}
            _ = v29
        default:
            if ((fi18 << ((18446744073709551615) & 63)) >= fi18) {
                for fi30 := 0; fi30 < 18; fi30++ {
                    v31 := GenS1{F0: 0}
                    _ = fi30
                    _ = v31
                }
                for fi32 := 0; fi32 < 1; fi32++ {
                    var v33 int8 = v4
                    if (!true) {
                        v17 += 1
                    }
                    var v34 uint = (^(v5))
                    v35 := GenS1{F0: -915}
                    if v1 {
                        fmt.Println(v10)
                        v36 := GenS1{F0: 0}
                        var v37 int32 = -257
                        _ = v36
                        _ = v37
                    } else {
                        var v38 uint64 = 4294967295
                        _ = v38
                    }
                    var v39 int32 = 0
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                    _ = v39
                }
            }
        }
        _ = fi18
    }
    fmt.Println(v1, v4, v5)
    v40 := GenS1{F0: 42}
    var v41 int = 32767
    for fi42 := 0; fi42 < 9; fi42++ {
        var v43 int32 = -2147483648
        _ = fi42
        _ = v43
    }
    switch v41 {
    case 4294967295:
        _ = len(v2)
    case 42:
        switch ("hello" + (("" + ("" + ("a" + ("\n\t" + "\n\t")))) + "")) {
        case "foo bar":
            var v44 string = "a"
            _ = v44
        default:
            for fi45 := 0; fi45 < 10; fi45++ {
                switch ("" + "foo bar") {
                case "a":
                    var v46 bool = (1e10 != 3.14159)
                    _ = v46
                case "hello":
                    for fi47 := 0; fi47 < 9; fi47++ {
                        v48 := GenS1{F0: 1}
                        v10 = (^(v10))
                        v17 -= (v17 << ((1) & 63))
                        _ = fi47
                        _ = v48
                    }
                case "":
                    v16 = v16
                }
                v16 = v10
                v41 += (v41 % (((v41 & 0)) | 1))
                _ = fi45
            }
        }
    case 363:
        var v49 uint8 = 127
        _ = v49
    }
    fmt.Println(v17)
    var v50 string = (("" + "hello") + "\n\t")
    var v51 string = "hello"
    var v52 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
