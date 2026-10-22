package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 int16, p1 uint32) string {
    var v4 bool = true
    var v5 int8 = 0
    p1 = (p1 << ((255) & 63))
    v6 := GenS1{F0: (p0 & (p0 >> ((1) & 63))), F1: -448}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return ("a" + ("" + "\n\t"))
}
func genFn1(p0 uint8, p1 bool, p2 string) string {
    v7 := GenS1{F0: 736, F1: 0}
    if p1 {
        for fi8 := 0; fi8 < 5; fi8++ {
            switch -9223372036854775808 {
            case -2147483648:
                if p1 {
                    v9 := []int{fi8, (fi8 + (-(fi8))), 0}
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        if (2 > 0) {
                            p1 = p1
                            v11 := []int{32767, fi8, fi8, fi10}
                            _ = v11
                        } else {
                            v12 := GenS1{F0: 42, F1: (fi8 ^ (fi8 >> ((97) & 63)))}
                            if ((p0 * 232) < 160) {
                                p1 = p1
                                fmt.Println(p1, p2, fi10)
                            } else {
                                p2 = "\n\t"
                                var v13 int8 = 54
                                _ = v13
                            }
                            _ = v12
                        }
                        _ = fi10
                    }
                    if true {
                        var v14 uint16 = 42
                        var v15 float32 = (0.1 - (1.5 - -2.25))
                        var v16 float32 = 0.1
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
            case 1:
                v17 := GenS1{F0: 32767, F1: (fi8 >> ((2) & 63))}
                _ = v17
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                p0 = (p0 - (p0 + p0))
                _ = fi18
            }
            _ = fi8
        }
        var v19 int16 = 42
        _ = v19
    }
    var v20 bool = ((-931 == -176) || (!((0 <= 42) && false)))
    v21 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v7
    return p2
}

func main() {
    var v1 uint64 = 2
    v2 := []int{1}
    v3 := []uint32{}
    var v22 string = ("hello" + ("\n\t" + "a"))
    for fi23 := 0; fi23 < 3; fi23++ {
        v24 := GenS1{F0: 42, F1: fi23}
        _ = fi23
        _ = v24
    }
    if (false && false) {
        v2 = append(v2, 4294967295)
        var v25 bool = (v22 >= ("a" + "hello"))
        v22 = v22
        _ = v25
    } else {
        switch v22 {
        case "hello":
            v26 := GenS2{F0: v22}
            _ = v26
        }
        var v27 uint32 = 42
        for fi28 := 0; fi28 < 14; fi28++ {
            for fi29 := 0; fi29 < 10; fi29++ {
                switch (v22 + "a") {
                case "\n\t":
                    var v30 uint64 = v1
                    _ = v30
                }
                var v31 bool = false
                _ = fi29
                _ = v31
            }
            _ = fi28
        }
        _ = v27
    }
    v32 := GenS3{F0: 2147483647, F1: v1}
    for fi33 := 0; fi33 < 3; fi33++ {
        v22 = v22
        if (!(1e10 > 1e10)) {
            v34 := GenS3{F0: -1, F1: (^(v1))}
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            _ = v34
        }
        if true {
            if (!false) {
                var v35 int8 = 1
                if true {
                    var v36 string = ("a" + (v22 + "hello"))
                    _ = v36
                } else {
                    v22 = v22
                }
                _ = v35
            } else {
                var v37 uint64 = (v1 - (v1 / (((v1 << ((4294967295) & 63))) | 1)))
                _ = v37
            }
            var v38 uint32 = 783
            _ = v38
        }
        _ = fi33
    }
    if false {
        v39 := GenS1{F0: 799, F1: 29}
        var v40 int = 504
        _ = v39
        _ = v40
    } else {
        var v41 bool = false
        v42 := GenS3{F0: -2147483648, F1: (v1 & v1)}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v41
        _ = v42
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch v22 {
    case "foo bar":
        if true {
            v22 = v22
        } else {
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
            var v43 float32 = 3.14159
            _ = v43
        }
    case "a":
        if false {
            var v44 int32 = -395
            v45 := GenS1{F0: 32767, F1: -1}
            v44 += -1
            _ = v44
            _ = v45
        }
    }
    var v46 uint8 = 56
    if (((float32(-2.25) - 3.14159) > (-2.25 / 1.0)) && false) {
        var v47 int32 = 971
        _ = v47
    } else {
        if false {
            v46 = 255
        } else {
            v48 := GenS4{F0: (v46 % ((v46) | 1)), F1: 0}
            var v49 uint = 42
            _ = v48
            _ = v49
        }
        v46 += v46
        v22 = (v22 + "a")
    }
    var v50 string = ""
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v46:", v46)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v46
    _ = v50
}
