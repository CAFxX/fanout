package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int16; F1 uint32; F2 uint32 }
func genFn0() int32 {
    var v4 int32 = 2147483647
    for fi5 := 0; fi5 < 4; fi5++ {
        v4 = 684
        var v6 int64 = 355
        if (1.5 > -2.25) {
            var v7 int = 32768
            var v8 float32 = float32(0.0)
            _ = v7
            _ = v8
        } else {
            var v9 int32 = (v4 << ((42) & 63))
            if true {
                for fi10 := 0; fi10 < 19; fi10++ {
                    v11 := GenS1{F0: v9, F1: 42, F2: (!(0 < 127))}
                    _ = fi10
                    _ = v11
                }
                switch (-(fi5)) {
                case 32768:
                    var v12 float64 = 0.1
                    _ = v12
                }
            }
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    _ = v4
    return v4
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]int)
    var v3 bool = (true && true)
    v13 := GenS1{F0: 0, F1: 65535, F2: false}
    v14 := GenS1{F0: genFn0(), F1: 18446744073709551615, F2: v3}
    for fi15 := 0; fi15 < 20; fi15++ {
        if v3 {
            v16 := GenS1{F0: -2147483648, F1: 1, F2: false}
            var v17 string = "\n\t"
            _ = v16
            _ = v17
        } else {
            if false {
                var v18 bool = v3
                var v19 uint = 255
                for fi20 := 0; fi20 < 10; fi20++ {
                    v1 = v1
                    if false {
                        v21 := GenS1{F0: -919, F1: 18446744073709551615, F2: v18}
                        var v22 uint64 = 1
                        switch "hello" {
                        case "hello":
                            v19 += v19
                        }
                        var v23 bool = true
                        v24 := GenS2{F0: v1}
                        v3 = (((^(v22)) != (^(v22))) && ((v22 * v22) < (v22 ^ 2)))
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = fi20
                }
                _ = v18
                _ = v19
            }
            if false {
                switch ("" + "hello") {
                case "a":
                    v1 += -123
                case "foo bar":
                    for fi25 := 0; fi25 < 9; fi25++ {
                        fmt.Println(v1, v1, v1)
                        v26 := GenS3{F0: "\n\t"}
                        v1 -= v1
                        _ = fi25
                        _ = v26
                    }
                case "hello":
                    if (1.5 <= (0.0 - 0.1)) {
                        for fi27 := 0; fi27 < 18; fi27++ {
                            var v28 uint32 = 695
                            v3 = ((v28 >> ((uint(v28)) & 63)) > 901)
                            _ = fi27
                            _ = v28
                        }
                    }
                }
            } else {
                var v29 float32 = 0.0
                v30 := GenS4{F0: (!("\n\t" <= ""))}
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 15; fi31++ {
                if v3 {
                    fmt.Println(v1, fi31, fi15)
                } else {
                    v32 := v2[(^(fi31))]
                    _ = v32
                }
                _ = fi31
            }
        }
        var v33 float64 = 1e10
        var v34 string = "a"
        _ = fi15
        _ = v33
        _ = v34
    }
    var v35 int8 = (v1 | v1)
    for fi36 := 0; fi36 < 7; fi36++ {
        delete(v2, (fi36 * (fi36 & (fi36 >> ((2) & 63)))))
        v1 -= 0
        _ = fi36
    }
    var v37 bool = v3
    v37 = v3
    var v38 float32 = (((1.5 + ((float32(float64((3.14159 / -2.25))) / 0.1) + 0.1)) * 0.1) * (3.14159 / float32((0.1 - 1e10))))
    v39 := v2[-2147483648]
    for fi40 := 0; fi40 < 12; fi40++ {
        switch "\n\t" {
        case "a":
            v41 := v2[v39]
            _ = v41
        case "\n\t":
            v42 := GenS5{F0: 743, F1: 153}
            _ = v42
        }
        var v43 int = 32768
        v35 -= v1
        _ = fi40
        _ = v43
    }
    var v44 uint16 = 630
    var v45 uint16 = 32767
    for fi46 := 0; fi46 < 4; fi46++ {
        v37 = (v37 || ((v3 && v3) && (((-880 != 1) && v3) || (1 != 303))))
        _ = fi46
    }
    v45 -= 242
    v47 := GenS6{F0: 32767, F1: 1000, F2: 42}
    v38 += ((1e10 - v38) * v38)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v39
    _ = v44
    _ = v45
    _ = v47
}
