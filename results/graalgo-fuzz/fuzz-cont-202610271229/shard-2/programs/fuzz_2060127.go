package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 int }
type GenS4 struct { F0 int32; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint }
func genFn0(p0 string) int {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            p0 = p0
            switch fi4 {
            case 2147483647:
                var v6 uint32 = 42
                _ = v6
            }
            _ = fi5
        }
        if false {
            p0 = p0
            switch (fi4 - fi4) {
            case 128:
                v7 := []int{}
                _ = v7
            case -9223372036854775808:
                v8 := []int{fi4}
                _ = v8
            case 488:
                var v9 uint64 = 1
                _ = v9
            }
        } else {
            if true {
                p0 = ""
            }
            if (!(!((-(fi4)) > (fi4 / ((fi4) | 1))))) {
                if (!true) {
                    var v10 bool = true
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 int8 = -1
                        var v13 uint64 = 280
                        if (30 < (v12 & v12)) {
                            v10 = v10
                            v13 = (v13 - 2)
                        }
                        v14 := []int{fi11, fi11, (fi4 << ((255) & 63)), fi4}
                        if (!(135 <= -9223372036854775808)) {
                            var v15 uint64 = (^(v13))
                            _ = v15
                        } else {
                            var v16 int16 = 1000
                            _ = v16
                        }
                        if len(v14) > 0 {
                            v14[0] = 127
                        }
                        switch 0 {
                        case 32767:
                            if len(v14) > 0 {
                                v14[0] = (fi11 & (-(fi11)))
                            }
                        case 0:
                            fmt.Println(fi4)
                        default:
                            fmt.Println(p0, fi11, fi11)
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    switch "foo bar" {
                    case "hello":
                        if (65535 == 2147483647) {
                            v17 := []int{(fi4 * fi4), fi4, -2147483648}
                            _ = v17
                        } else {
                            v18 := []int{}
                            _ = v18
                        }
                    }
                    _ = v10
                } else {
                    p0 = p0
                    var v19 int16 = 32767
                    for fi20 := 0; fi20 < 10; fi20++ {
                        if false {
                            v19 -= v19
                            fmt.Println(v19, fi20, p0)
                            v21 := []int{}
                            _ = v21
                        }
                        switch 678 {
                        case -855:
                            v22 := []int{}
                            _ = v22
                        case 127:
                            fmt.Println(v19, p0)
                        case 32767:
                            var v23 float64 = (0.1 + 0.0)
                            _ = v23
                        }
                        _ = fi20
                    }
                    _ = v19
                }
            } else {
                var v24 uint8 = 2
                _ = v24
            }
        }
        var v25 bool = true
        _ = fi4
        _ = v25
    }
    p0 = p0
    _ = p0
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 bool = true
    if true {
        delete(v2, ("\n\t" + "hello"))
        v26 := GenS1{F0: (3.14159 != float64(3.14159))}
        _ = v26
    } else {
        fmt.Println(v3)
        for fi27 := 0; fi27 < 5; fi27++ {
            switch "" {
            case "foo bar":
                var v28 int = fi27
                _ = v28
            }
            switch "a" {
            case "\n\t":
                var v29 uint64 = 42
                _ = v29
            default:
                for fi30 := 0; fi30 < 4; fi30++ {
                    var v31 float64 = 1.5
                    var v32 int64 = int64(fi27)
                    delete(v2, "")
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
            for fi33 := 0; fi33 < 20; fi33++ {
                v3 = v3
                _ = fi33
            }
            _ = fi27
        }
        v1[2147483648] = ""
    }
    v3 = v3
    v34 := GenS2{F0: 506}
    if v3 {
        switch ("\n\t" + ("\n\t" + ("\n\t" + "hello"))) {
        case "hello":
            if v3 {
                var v35 bool = (-4 >= 255)
                v36 := []int{genFn0(("foo bar" + (("hello" + "a") + "a")))}
                for fi37 := 0; fi37 < 6; fi37++ {
                    v38 := GenS3{F0: (-(fi37)), F1: 65535, F2: fi37}
                    var v39 string = ""
                    v39 = ((v39 + (((v39 + ((v39 + v39) + "a")) + ((v39 + "a") + (v39 + ""))) + (((v39 + v39) + v39) + ("a" + "")))) + (((v39 + (v39 + "foo bar")) + (v39 + "\n\t")) + "\n\t"))
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                _ = v35
                _ = v36
            }
        case "a":
            var v40 int = 1
            _ = v40
        case "":
            var v41 bool = ((1e10 + float32(1.5)) != 1.5)
            _ = v41
        }
    }
    v42 := GenS4{F0: 42, F1: 865, F2: -776}
    for fi43 := 0; fi43 < 17; fi43++ {
        delete(v2, "foo bar")
        v3 = ((1000 >= 32767) && true)
        fmt.Println(v3, fi43)
        _ = fi43
    }
    var v44 int64 = 9223372036854775807
    v45 := GenS5{F0: 1, F1: 4294967295}
    fmt.Println(v3, v3)
    fmt.Println(v44, v3)
    for fi46 := 0; fi46 < 10; fi46++ {
        delete(v2, "hello")
        var v47 float32 = float32((float64(float32(((float64(3.14159) - 3.14159) - 1.5))) * 1.5))
        var v48 uint16 = 0
        _ = fi46
        _ = v47
        _ = v48
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        v44 += (-(v44))
        fmt.Println(fi49, v3)
        switch (fi49 + -623) {
        case 128:
            v50 := GenS6{F0: 18446744073709551615}
            _ = v50
        case 2147483647:
            for fi51 := 0; fi51 < 7; fi51++ {
                var v52 float32 = (0.1 / -2.25)
                var v53 uint16 = 65535
                var v54 uint = 65535
                _ = fi51
                _ = v52
                _ = v53
                _ = v54
            }
        case -845:
            for fi55 := 0; fi55 < 7; fi55++ {
                v56 := GenS2{F0: 32767}
                v57 := []int{-2147483648, fi49, -235, -9223372036854775808}
                _ = fi55
                _ = v56
                _ = v57
            }
        }
        _ = fi49
    }
    var v58 int32 = -937
    v3 = (42 != 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v42
    _ = v44
    _ = v45
    _ = v58
}
