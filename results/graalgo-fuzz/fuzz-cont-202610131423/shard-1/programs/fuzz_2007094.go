package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint64; F2 bool }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint64; F1 int8; F2 uint64 }

func main() {
    var v1 uint16 = 65535
    var v2 float64 = ((-2.25 - ((1e10 / 1.5) / 3.14159)) - -2.25)
    var v3 uint8 = 91
    switch 32767 {
    case 128:
        for fi4 := 0; fi4 < 5; fi4++ {
            v5 := []int{fi4, (fi4 / (((-(fi4))) | 1))}
            _ = fi4
            _ = v5
        }
    }
    fmt.Println(v3, v1, v2)
    for fi6 := 0; fi6 < 9; fi6++ {
        v3 += v3
        _ = fi6
    }
    if false {
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                v2 = 0.0
                _ = fi8
            }
            v9 := []int{(^(fi7))}
            _ = fi7
            _ = v9
        }
        v2 += v2
    } else {
        v10 := GenS1{F0: 287}
        _ = v10
    }
    v11 := []int{42, 32768, -2147483648, -2147483648}
    var v12 string = "foo bar"
    var v13 uint16 = (v1 % ((v1) | 1))
    if true {
        if false {
            if false {
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS1{F0: 709}
                    fmt.Println(v13, fi14)
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 int8 = int8(fi16)
                    _ = fi16
                    _ = v17
                }
                var v18 uint64 = 4294967295
                _ = v18
            }
            v11 = append(v11, 461)
        } else {
            var v19 uint8 = 4
            _ = v19
        }
        if true {
            fmt.Println(v2, v2)
        }
        v2 -= 1.5
    }
    var v20 int = -9223372036854775808
    if false {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 uint16 = (v1 >> ((847) & 63))
            _ = fi21
            _ = v22
        }
        switch (v12 + v12) {
        case "a":
            if (-991 == 0) {
                v11 = append(v11, v20)
                v3 = v3
            } else {
                if (32767 > 1000) {
                    if (((!false) && (false || (255 == (v1 & uint16(v3))))) && (-88 == 42)) {
                        v12 = ((v12 + "a") + "")
                        fmt.Println(v20, v1, v12)
                        if (2147483647 > 2147483647) {
                            v23 := GenS2{F0: true}
                            v24 := GenS3{F0: -86, F1: 289, F2: false}
                            if false {
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    _ = len(v11)
                                    var v26 uint64 = 4294967295
                                    _ = fi25
                                    _ = v26
                                }
                                v27 := GenS3{F0: 127, F1: 2, F2: false}
                                v1 -= v1
                                _ = v27
                            }
                            _ = v23
                            _ = v24
                        } else {
                            v12 = ("\n\t" + (("foo bar" + "") + v12))
                        }
                    } else {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        for fi28 := 0; fi28 < 13; fi28++ {
                            var v29 string = v12
                            if false {
                                v3 += (v3 % (((v3 % ((v3) | 1))) | 1))
                                var v30 int32 = 0
                                _ = v30
                            }
                            _ = fi28
                            _ = v29
                        }
                    }
                }
                v20 = (^(v20))
            }
        case "":
            v13 = (v1 << ((980) & 63))
        }
        v31 := GenS4{F0: (!(false || false)), F1: 255}
        _ = v31
    } else {
        var v32 uint8 = v3
        switch (-(v20)) {
        case 32768:
            switch ((((v12 + (v12 + ("" + "hello"))) + v12) + (v12 + "")) + v12) {
            case "hello":
                if false {
                    v1 = v13
                    fmt.Println(v3)
                    v33 := GenS5{F0: 1, F1: 22, F2: 18446744073709551615}
                    _ = v33
                } else {
                    fmt.Println(v32, v20, v2)
                }
            }
        case 0:
            for fi34 := 0; fi34 < 20; fi34++ {
                v35 := GenS3{F0: 100, F1: 18446744073709551615, F2: true}
                _ = fi34
                _ = v35
            }
        }
        _ = v32
    }
    if len(v11) > 0 {
        v11[0] = (v20 << ((4294967295) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
}
