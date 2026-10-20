package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint64 = 1
    var v2 int8 = 0
    v3 := make(map[uint32]uint32)
    var v4 int64 = 0
    v5 := make(map[string]int)
    v6 := make(map[int]string)
    switch "foo bar" {
    case "hello":
        if true {
            fmt.Println(v4, v1, v1)
        } else {
            v2 -= -109
            for fi7 := 0; fi7 < 20; fi7++ {
                switch ("a" + "hello") {
                case "":
                    v8 := []int{fi7, (fi7 | 32768)}
                    _ = v8
                default:
                    fmt.Println(v4)
                }
                _ = fi7
            }
            var v9 float32 = 0.1
            _ = v9
        }
    case "foo bar":
        v4 += 42
    default:
        v10 := GenS1{F0: (v1 >> ((255) & 63))}
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        switch 42 {
        case -2147483648:
            v12 := []int{}
            _ = v12
        case 0:
            v13 := []int{(fi11 ^ fi11), (-(fi11)), fi11, fi11}
            _ = v13
        case -86:
            v2 = int8(v1)
        }
        var v14 bool = (-9223372036854775808 == v4)
        _ = fi11
        _ = v14
    }
    v2 = v2
    switch 547 {
    case 9223372036854775807:
        var v15 bool = true
        _ = v15
    case 206:
        var v16 int32 = int32(v4)
        _ = v16
    }
    if true {
        var v17 uint32 = 2
        _ = v17
    } else {
        fmt.Println(v2, v2, v1)
        for fi18 := 0; fi18 < 3; fi18++ {
            if false {
                var v19 uint16 = 42
                if true {
                    v20 := []int{(fi18 - fi18)}
                    var v21 int32 = -501
                    v22 := GenS1{F0: v1}
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    v4 += v4
                    if (!false) {
                        v1 += (v1 / ((1) | 1))
                        for fi23 := 0; fi23 < 11; fi23++ {
                            var v24 bool = true
                            _ = fi23
                            _ = v24
                        }
                    }
                }
                if (42 != 16) {
                    fmt.Println(v19, v19)
                    v1 -= v1
                }
                _ = v19
            }
            v25 := []int{fi18, 14, 127, fi18}
            if len(v25) > 0 {
                v25[0] = (fi18 << ((0) & 63))
            }
            _ = fi18
            _ = v25
        }
        v4 += 9223372036854775807
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        if (2147483648 >= 2147483648) {
            v27 := GenS2{F0: 478}
            if ((v1 / (((v1 * v1)) | 1)) > (v1 + v1)) {
                v1 = 4294967295
                for fi28 := 0; fi28 < 19; fi28++ {
                    if (!(1e10 >= (1.5 * float32(0.1)))) {
                        v2 = (v2 % ((v2) | 1))
                        switch (fi26 | -9223372036854775808) {
                        case -2147483648:
                            v1 = 0
                        default:
                            var v29 float32 = 1.5
                            _ = v29
                        }
                    } else {
                        var v30 uint16 = 4
                        v1 -= (v1 * (v1 * v1))
                        var v31 uint32 = 65535
                        _ = v30
                        _ = v31
                    }
                    fmt.Println(fi26)
                    var v32 float64 = 1e10
                    _ = fi28
                    _ = v32
                }
            }
            _ = v27
        } else {
            v33 := GenS1{F0: v1}
            if ((v4 == v4) || (768 > 4294967295)) {
                var v34 int = (fi26 >> ((915) & 63))
                var v35 uint32 = 720
                for fi36 := 0; fi36 < 9; fi36++ {
                    v35 = (v35 << ((42) & 63))
                    fmt.Println(v4)
                    v34 = (v34 - fi26)
                    _ = fi36
                }
                _ = v34
                _ = v35
            }
            _ = v33
        }
        v37 := GenS3{F0: 450}
        _ = fi26
        _ = v37
    }
    switch "a" {
    case "a":
        v4 = 42
    case "\n\t":
        var v38 uint16 = 32767
        _ = v38
    default:
        if ((v2 >> ((255) & 63)) == 1) {
            for fi39 := 0; fi39 < 15; fi39++ {
                var v40 uint = 18446744073709551615
                for fi41 := 0; fi41 < 2; fi41++ {
                    fmt.Println(fi41, v40, v2)
                    _ = fi41
                }
                v1 = v1
                _ = fi39
                _ = v40
            }
        }
    }
    v1 = (v1 | (^(v1)))
    if (!(255 > 0)) {
        for fi42 := 0; fi42 < 17; fi42++ {
            fmt.Println(v4)
            _ = fi42
        }
        v4 += (v4 * (v4 << ((255) & 63)))
    }
    for fi43 := 0; fi43 < 14; fi43++ {
        var v44 string = "a"
        _ = fi43
        _ = v44
    }
    if (-105 != 680) {
        v45 := v6[2147483647]
        var v46 bool = false
        var v47 uint8 = 49
        _ = v45
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
