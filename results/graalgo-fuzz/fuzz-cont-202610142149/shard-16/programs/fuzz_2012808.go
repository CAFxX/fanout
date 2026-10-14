package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint) uint32 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int = fi7
        var v9 int8 = 127
        p0 = p0
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 += (p0 / ((p0) | 1))
    if true {
        var v10 uint64 = 844
        v11 := []int{4294967295}
        fmt.Println(p0, v10)
        _ = v10
        _ = v11
    } else {
        var v12 float64 = 1.5
        _ = v12
    }
    fmt.Println(p0)
    var v13 int16 = -985
    _ = p0
    _ = v13
    return 4294967295
}

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 65
    v3 := []uint32{(v2 >> ((42) & 63)), (^(v2)), 42}
    var v4 int = -807
    var v5 string = (("a" + "hello") + "hello")
    var v6 uint64 = 0
    var v14 float64 = 0.0
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS1{F0: (v6 / (((^(v6))) | 1))}
        _ = len(v3)
        _ = fi15
        _ = v16
    }
    fmt.Println(v14, v5)
    var v17 uint64 = v6
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int8 = -1
        _ = fi18
        _ = v19
    }
    fmt.Println(v4, v17, v1)
    v20 := GenS1{F0: v17}
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 string = "foo bar"
        _ = fi21
        _ = v22
    }
    if true {
        var v23 float64 = v14
        _ = v23
    } else {
        var v24 int16 = 0
        v25 := GenS2{F0: int16(v6)}
        if false {
            var v26 uint16 = 0
            if ((!(v26 > v26)) && false) {
                v3 = append(v3, 1)
                v17 += 4294967295
                for fi27 := 0; fi27 < 20; fi27++ {
                    switch "\n\t" {
                    case "foo bar":
                        v28 := GenS1{F0: 4294967295}
                        _ = v28
                    case "\n\t":
                        fmt.Println(v17, v26)
                    case "":
                        var v29 bool = (!false)
                        _ = v29
                    default:
                        switch fi27 {
                        case 128:
                            v30 := GenS2{F0: v24}
                            _ = v30
                        case -827:
                            if ((false && true) || false) {
                                v3 = append(v3, (^(v2)))
                            } else {
                                var v31 uint32 = v2
                                _ = v31
                            }
                        default:
                            var v32 float64 = 0.1
                            _ = v32
                        }
                    }
                    fmt.Println(v1)
                    v4 += v4
                    _ = fi27
                }
            } else {
                if false {
                    v26 += v26
                    fmt.Println(v26)
                } else {
                    var v33 int16 = (v24 + (v24 * (v24 - (v24 & v24))))
                    var v34 int16 = v24
                    _ = v33
                    _ = v34
                }
            }
            if len(v3) > 0 {
                v3[0] = v2
            }
            _ = v26
        } else {
            if (32767 != 32767) {
                if false {
                    for fi35 := 0; fi35 < 20; fi35++ {
                        v2 += (^(v2))
                        for fi36 := 0; fi36 < 14; fi36++ {
                            v17 += v6
                            _ = fi36
                        }
                        v2 = v2
                        _ = fi35
                    }
                }
                var v37 uint32 = (^(v2))
                _ = v37
            }
            var v38 uint8 = 128
            var v39 int64 = 187
            _ = v38
            _ = v39
        }
        _ = v24
        _ = v25
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        v5 = (v5 + v5)
        switch "foo bar" {
        case "hello":
            var v41 uint = 4294967295
            _ = v41
        case "\n\t":
            var v42 int64 = 815
            _ = v42
        case "a":
            switch (("a" + v5) + "a") {
            case "a":
                var v43 uint8 = 2
                _ = v43
            }
        default:
            v14 += float64(-2.25)
        }
        _ = fi40
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
