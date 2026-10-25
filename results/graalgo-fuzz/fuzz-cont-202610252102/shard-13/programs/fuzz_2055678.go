package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 uint64; F2 int }
func genFn0(p0 uint64) int16 {
    switch "a" {
    case "":
        p0 = (p0 | p0)
    default:
        if (false && true) {
            v4 := []int{}
            p0 = 2
            _ = v4
        }
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        switch "foo bar" {
        case "":
            for fi6 := 0; fi6 < 1; fi6++ {
                v7 := GenS1{F0: 2147483647}
                _ = fi6
                _ = v7
            }
        case "a":
            var v8 int = fi5
            _ = v8
        default:
            p0 += 42
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{fi5}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int8 = -1
            var v13 int = (fi5 << ((42) & 63))
            v14 := []int{32768, 32767, (v13 << ((625) & 63))}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi5
    }
    _ = p0
    return -1
}

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 uint16 = 255
    v3 := []uint32{726, 0, 2}
    var v15 int8 = 100
    var v16 uint = 2
    if false {
        if false {
            if true {
                v15 += v15
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v17 uint32 = 2147483648
                _ = v17
            } else {
                for fi18 := 0; fi18 < 2; fi18++ {
                    v19 := GenS2{F0: 1, F1: ("foo bar" + v1)}
                    v1 = v1
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                    _ = fi18
                    _ = v19
                }
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            v20 := GenS2{F0: 42, F1: v1}
            _ = v20
        }
        fmt.Println(v16)
    } else {
        var v21 float64 = 1.5
        var v22 int16 = 1000
        _ = v21
        _ = v22
    }
    var v23 int16 = 32767
    v2 = 65535
    switch ((("hello" + v1) + v1) + "\n\t") {
    case "a":
        switch -1 {
        case 0:
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 float64 = 1.5
                _ = fi24
                _ = v25
            }
        case 128:
            v1 = v1
        case -2147483648:
            if true {
                fmt.Println(v2)
                for fi26 := 0; fi26 < 11; fi26++ {
                    v16 = (v16 / ((v16) | 1))
                    for fi27 := 0; fi27 < 2; fi27++ {
                        v28 := GenS3{F0: 203, F1: 1}
                        _ = fi27
                        _ = v28
                    }
                    var v29 int8 = (v15 ^ 127)
                    _ = fi26
                    _ = v29
                }
            } else {
                v30 := GenS3{F0: 2, F1: 156}
                v15 = 127
                _ = v30
            }
        }
    case "":
        var v31 float64 = float64(3.14159)
        _ = v31
    default:
        for fi32 := 0; fi32 < 8; fi32++ {
            if (!(906 >= 0)) {
                if false {
                    switch (fi32 << ((v16) & 63)) {
                    case -921:
                        v33 := GenS4{F0: 99, F1: 27, F2: 42}
                        _ = v33
                    case -2147483648:
                        if (genFn0(565) == v23) {
                            var v34 uint8 = 43
                            for fi35 := 0; fi35 < 2; fi35++ {
                                v36 := GenS5{F0: (((v1 + v1) + ((("\n\t" + v1) + "\n\t") + ("\n\t" + (v1 + (v1 + v1))))) + v1)}
                                v34 += 179
                                v23 += v23
                                v16 += (^(v16))
                                _ = fi35
                                _ = v36
                            }
                            var v37 int64 = 0
                            _ = v34
                            _ = v37
                        }
                    }
                    for fi38 := 0; fi38 < 16; fi38++ {
                        var v39 bool = true
                        var v40 string = v1
                        var v41 float64 = -2.25
                        v40 = v1
                        var v42 int = (fi38 << ((v16) & 63))
                        _ = len(v3)
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                }
                if (!(v23 <= v23)) {
                    v3 = append(v3, 2147483647)
                } else {
                    fmt.Println(v16)
                    var v43 int32 = 227
                    var v44 float32 = 1.5
                    _ = v43
                    _ = v44
                }
                var v45 string = "foo bar"
                _ = v45
            } else {
                v46 := GenS6{F0: v23, F1: 0, F2: (-(fi32))}
                if (v23 == 0) {
                    fmt.Println(fi32, v2)
                }
                _ = v46
            }
            _ = fi32
        }
    }
    if len(v3) > 0 {
        v3[0] = 915
    }
    var v47 uint = v16
    switch -9223372036854775808 {
    case 2147483647:
        fmt.Println(v16, v2, v1)
    case 4294967295:
        var v48 bool = (!false)
        _ = v48
    case 9223372036854775807:
        if len(v3) > 0 {
            v3[0] = uint32(v2)
        }
    default:
        var v49 float64 = -2.25
        _ = v49
    }
    v47 += v47
    switch 0 {
    case 42:
        for fi50 := 0; fi50 < 17; fi50++ {
            var v51 bool = false
            for fi52 := 0; fi52 < 16; fi52++ {
                v51 = (true && (v1 == (("" + "\n\t") + v1)))
                v3 = append(v3, 415)
                _ = fi52
            }
            var v53 uint64 = 1
            _ = fi50
            _ = v51
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v47
}
