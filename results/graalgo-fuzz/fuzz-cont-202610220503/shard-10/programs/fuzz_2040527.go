package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int }
func genFn0(p0 int64, p1 int32) int8 {
    if false {
        var v5 uint = 421
        if true {
            var v6 int16 = 0
            for fi7 := 0; fi7 < 2; fi7++ {
                var v8 int64 = (p0 - p0)
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        _ = v5
    }
    if true {
        v9 := []int{4294967295, -1, 720}
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                if (!true) {
                    v12 := GenS1{F0: fi10}
                    _ = v12
                }
                var v13 int8 = 1
                _ = fi11
                _ = v13
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi10
        }
        if false {
            v14 := GenS1{F0: 1}
            var v15 int32 = (p1 >> ((42) & 63))
            switch -14 {
            case 0:
                fmt.Println(v15, p0, v15)
            }
            _ = v14
            _ = v15
        }
        _ = v9
    } else {
        fmt.Println(p1, p0)
        switch 32768 {
        case 32768:
            var v16 int64 = 42
            _ = v16
        case 32767:
            for fi17 := 0; fi17 < 12; fi17++ {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := []int{fi18, fi18, fi17}
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        }
    }
    var v20 int = -678
    if true {
        if false {
            fmt.Println(v20, v20, p0)
        }
        fmt.Println(p0)
    } else {
        if false {
            var v21 uint = 930
            v20 -= 967
            var v22 int64 = p0
            _ = v21
            _ = v22
        }
    }
    var v23 bool = false
    _ = p0
    _ = p1
    _ = v20
    _ = v23
    return 122
}

func main() {
    v1 := []string{}
    var v2 uint16 = 1
    var v3 int16 = 42
    var v4 string = ""
    if ("\n\t" >= "") {
        v24 := GenS1{F0: 10}
        v25 := GenS1{F0: 0}
        var v26 string = (v4 + v4)
        _ = v24
        _ = v25
        _ = v26
    } else {
        if (3.14159 <= 0.1) {
            v27 := GenS2{F0: 685, F1: v2}
            var v28 float64 = (0.0 / 1.5)
            _ = v27
            _ = v28
        }
        fmt.Println(v3)
    }
    var v29 int = 32767
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 string = ""
        switch v4 {
        case "\n\t":
            for fi32 := 0; fi32 < 13; fi32++ {
                v3 += (v3 ^ v3)
                v33 := GenS3{F0: 0}
                _ = fi32
                _ = v33
            }
        case "":
            v34 := GenS4{F0: 2, F1: 1}
            _ = v34
        case "hello":
            var v35 bool = (false || false)
            _ = v35
        }
        _ = fi30
        _ = v31
    }
    if false {
        for fi36 := 0; fi36 < 14; fi36++ {
            v2 += v2
            var v37 uint8 = 1
            for fi38 := 0; fi38 < 12; fi38++ {
                var v39 uint64 = 4294967295
                if true {
                    if false {
                        v40 := GenS1{F0: fi38}
                        _ = v40
                    }
                } else {
                    v29 = (fi38 + (fi38 >> ((1) & 63)))
                    v4 = (v4 + "a")
                    var v41 uint32 = 849
                    _ = v41
                }
                _ = fi38
                _ = v39
            }
            _ = fi36
            _ = v37
        }
    } else {
        var v42 string = v4
        switch -1 {
        case 268:
            fmt.Println(v2, v2, v2)
        case 793:
            var v43 uint16 = (v2 + v2)
            _ = v43
        }
        _ = v42
    }
    var v44 uint32 = 65535
    v45 := GenS5{F0: v29}
    switch (v4 + "foo bar") {
    case "":
        var v46 uint16 = (v2 - 72)
        _ = v46
    case "hello":
        if len(v1) > 0 {
            v1[0] = (((v4 + v4) + (v4 + v4)) + v4)
        }
    }
    switch 758 {
    case 32767:
        for fi47 := 0; fi47 < 14; fi47++ {
            v29 -= -9223372036854775808
            for fi48 := 0; fi48 < 10; fi48++ {
                v29 += (fi47 - fi47)
                v3 = (v3 & 42)
                v3 = v3
                _ = fi48
            }
            v29 += v29
            _ = fi47
        }
    default:
        switch v4 {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "hello":
            v1 = append(v1, v4)
        case "\n\t":
            for fi49 := 0; fi49 < 17; fi49++ {
                v2 += (v2 >> ((1) & 63))
                v29 = fi49
                for fi50 := 0; fi50 < 20; fi50++ {
                    switch fi50 {
                    case -2147483648:
                        if true {
                            fmt.Println(fi50, v29)
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                    _ = fi50
                }
                _ = fi49
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v45
}
