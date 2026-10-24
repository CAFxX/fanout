package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint16; F2 bool }
type GenS7 struct { F0 uint16; F1 uint8; F2 uint }

func main() {
    var v1 uint16 = 65535
    var v2 float32 = 1.5
    var v3 string = "a"
    var v4 float64 = 0.1
    for fi5 := 0; fi5 < 1; fi5++ {
        if false {
            v4 = v4
        } else {
            v1 = (v1 - (v1 / (((v1 << ((1) & 63))) | 1)))
            var v6 string = "hello"
            v7 := GenS1{F0: 501, F1: -1}
            _ = v6
            _ = v7
        }
        var v8 int8 = 102
        var v9 int = (fi5 >> ((927) & 63))
        _ = fi5
        _ = v8
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS2{F0: v1, F1: 702, F2: v1}
            _ = fi10
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := []int{42}
                v4 = (1.5 * -2.25)
                _ = fi13
                _ = v14
            }
            if false {
                var v15 bool = true
                v16 := GenS3{F0: 732, F1: 125}
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        v17 := GenS3{F0: 199, F1: -1}
        _ = v17
    }
    v1 = 0
    v18 := []int{-1}
    var v19 string = ("\n\t" + v3)
    var v20 int = -2147483648
    v4 += v4
    v19 = ("a" + (v19 + "\n\t"))
    v1 -= v1
    v4 = float64(((-2.25 - float32((3.14159 + v4))) / v2))
    if false {
        fmt.Println(v2)
    }
    _ = len(v18)
    if false {
        var v21 float32 = v2
        v22 := GenS3{F0: 539, F1: 0}
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 15; fi23++ {
            v18 = append(v18, (fi23 * fi23))
            _ = fi23
        }
        switch (v3 + (((v19 + ("" + "")) + v3) + ("\n\t" + ""))) {
        case "a":
            for fi24 := 0; fi24 < 18; fi24++ {
                v19 = v19
                _ = len(v18)
                fmt.Println(v3)
                _ = fi24
            }
        case "hello":
            for fi25 := 0; fi25 < 16; fi25++ {
                if len(v18) > 0 {
                    v18[0] = (v20 * (v20 << ((255) & 63)))
                }
                var v26 uint16 = (v1 >> ((4294967295) & 63))
                v2 = (v2 + v2)
                _ = fi25
                _ = v26
            }
        }
        if true {
            v27 := GenS4{F0: true, F1: v1, F2: 42}
            _ = v27
        }
    }
    if true {
        v28 := GenS1{F0: -2147483648, F1: -1}
        _ = v28
    } else {
        var v29 uint16 = 813
        if true {
            if true {
                v1 = v29
                var v30 int8 = 42
                _ = v30
            } else {
                if len(v18) > 0 {
                    v18[0] = (^(v20))
                }
                for fi31 := 0; fi31 < 20; fi31++ {
                    v32 := GenS5{F0: -9223372036854775808}
                    _ = fi31
                    _ = v32
                }
                v2 = v2
            }
        } else {
            for fi33 := 0; fi33 < 15; fi33++ {
                v34 := GenS6{F0: (-(fi33)), F1: 0, F2: true}
                for fi35 := 0; fi35 < 7; fi35++ {
                    v1 -= v29
                    v36 := GenS3{F0: 1000, F1: 100}
                    switch 511 {
                    case 4294967295:
                        v37 := GenS5{F0: 0}
                        _ = v37
                    }
                    _ = fi35
                    _ = v36
                }
                if (v2 <= v2) {
                    v19 = v19
                }
                _ = fi33
                _ = v34
            }
        }
        if false {
            for fi38 := 0; fi38 < 1; fi38++ {
                v39 := GenS7{F0: v1, F1: 2, F2: 923}
                _ = fi38
                _ = v39
            }
            v40 := GenS2{F0: (v1 - (^(v1))), F1: 42, F2: 88}
            _ = v40
        } else {
            v18 = append(v18, 128)
            if (191 >= (v29 << ((747) & 63))) {
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                if (!(v20 >= (v20 | v20))) {
                    v20 += v20
                    v2 = (v2 + v2)
                    fmt.Println(v19, v19)
                }
            } else {
                for fi41 := 0; fi41 < 19; fi41++ {
                    var v42 uint32 = 2147483648
                    v43 := GenS5{F0: 9223372036854775807}
                    if true {
                        var v44 float32 = v2
                        var v45 bool = true
                        fmt.Println(v1, v29, v20)
                        _ = v44
                        _ = v45
                    } else {
                        var v46 uint64 = 2
                        _ = v46
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
        }
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
