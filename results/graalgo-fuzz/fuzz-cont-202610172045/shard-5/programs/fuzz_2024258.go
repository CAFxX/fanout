package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int16; F1 bool }
func genFn0(p0 bool, p1 string, p2 uint8) int64 {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 uint = 1
        p1 = ((p1 + p1) + p1)
        p1 = p1
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 int8, p1 uint, p2 int) int16 {
    v7 := []int{(p2 - -936), (p2 * p2), p2}
    var v8 bool = (42 >= 380)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = -982
    var v3 uint32 = 2
    v4 := make(map[int]string)
    var v9 uint = 668
    if (!true) {
        var v10 int = -1
        _ = v10
    } else {
        fmt.Println(v2, v9)
        v11 := []int{-9223372036854775808, 9223372036854775807}
        v3 = 671
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: false, F1: true, F2: 1}
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: (42 > 255), F1: (1 < 527), F2: -352}
    v15 := GenS1{F0: false, F1: false, F2: -2147483648}
    var v16 float32 = (1e10 + float32(0.0))
    if false {
        v16 = v16
        fmt.Println(v2, v2, v2)
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            switch (fi17 / ((fi17) | 1)) {
            case -1:
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 int8 = 1
                    _ = fi18
                    _ = v19
                }
            case 9223372036854775807:
                switch (("" + "a") + (("\n\t" + "") + "a")) {
                case "\n\t":
                    if true {
                        switch "foo bar" {
                        case "":
                            v9 -= 112
                        case "a":
                            fmt.Println(v3)
                        }
                    }
                case "":
                    var v20 int = -109
                    _ = v20
                }
            case 1:
                v3 = v3
            }
            _ = fi17
        }
        var v21 uint8 = 2
        v3 += (v3 << ((517) & 63))
        _ = v21
    }
    var v22 float32 = v16
    var v23 int = 409
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 10; fi25++ {
            v9 += v9
            _ = fi25
        }
        _ = fi24
    }
    v2 += -390
    switch "a" {
    case "foo bar":
        var v26 uint64 = 444
        _ = v26
    }
    if true {
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 int32 = 164
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 bool = (uint8(v3) == 2)
            if v30 {
                delete(v4, (-(v23)))
                if ((3.14159 > ((3.14159 + float64(v16)) + (0.0 * 3.14159))) && (v22 < ((v16 + (v16 * -2.25)) + float32(0.0)))) {
                    for fi31 := 0; fi31 < 6; fi31++ {
                        var v32 string = ("a" + "hello")
                        _ = fi31
                        _ = v32
                    }
                    if ((78 > -9223372036854775808) && true) {
                        if (!true) {
                            fmt.Println(v22, v30)
                        }
                        v16 = float32((0.1 + ((1e10 * 0.1) - -2.25)))
                        var v33 int = (v23 & v23)
                        _ = v33
                    }
                    v34 := GenS2{F0: 127, F1: "\n\t", F2: 42}
                    _ = v34
                }
            } else {
                v23 = int(v9)
                var v35 int32 = -2147483648
                fmt.Println(v16, fi29)
                _ = v35
            }
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 int32 = -2147483648
                var v38 int64 = 0
                var v39 float64 = 1.5
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
            _ = fi29
            _ = v30
        }
        for fi40 := 0; fi40 < 18; fi40++ {
            for fi41 := 0; fi41 < 16; fi41++ {
                if false {
                    v3 = (v3 >> ((2) & 63))
                    fmt.Println(v3, v3, v23)
                } else {
                    _ = len(v4)
                }
                if ("a" >= "a") {
                    var v42 uint8 = 127
                    if true {
                        v2 = (v2 % ((32767) | 1))
                        switch "" {
                        case "a":
                            var v43 uint64 = 455
                            _ = v43
                        }
                        v16 += 3.14159
                    }
                    v44 := GenS2{F0: -1, F1: "a", F2: 0}
                    v45 := GenS3{F0: -1, F1: int(v42)}
                    _ = v42
                    _ = v44
                    _ = v45
                }
                fmt.Println(fi40, fi41, v9)
                _ = fi41
            }
            v46 := GenS4{F0: (v2 >> ((v9) & 63)), F1: true}
            _ = fi40
            _ = v46
        }
    } else {
        v47 := []int{}
        _ = v47
    }
    var v48 int64 = -9223372036854775808
    fmt.Println(v23)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v48
    _ = v9
}
