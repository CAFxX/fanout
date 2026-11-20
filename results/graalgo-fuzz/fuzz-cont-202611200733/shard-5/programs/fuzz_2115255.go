package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 bool; F1 bool; F2 int16 }
func genFn0(p0 int) int64 {
    p0 = (^(p0))
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := GenS1{F0: int64(p0)}
        _ = fi4
        _ = v5
    }
    switch 128 {
    case 1:
        var v6 int8 = -1
        _ = v6
    case -1:
        switch ("\n\t" + "\n\t") {
        case "":
            var v7 uint16 = 0
            _ = v7
        }
    case 671:
        p0 = -9223372036854775808
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = -47
    var v3 uint32 = 42
    v8 := GenS2{F0: -1, F1: 2}
    if (-2.25 >= (3.14159 * (1.5 + 1.5))) {
        if (false && true) {
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := v1["\n\t"]
                fmt.Println(v10, fi9, v2)
                v11 := []int{}
                _ = fi9
                _ = v10
                _ = v11
            }
            v12 := GenS3{F0: (v3 % (((^(v3))) | 1))}
            if false {
                switch (v2 >> ((255) & 63)) {
                case 42:
                    var v13 uint8 = 57
                    _ = v13
                case -184:
                    v14 := GenS4{F0: v3, F1: 1}
                    _ = v14
                }
                var v15 float32 = 1e10
                _ = v15
            } else {
                var v16 uint64 = 186
                var v17 int16 = -1
                if (!false) {
                    var v18 uint64 = (v16 >> ((2) & 63))
                    _ = v18
                }
                _ = v16
                _ = v17
            }
            _ = v12
        }
    } else {
        var v19 uint = 325
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int64 = 1
        _ = fi20
        _ = v21
    }
    if false {
        fmt.Println(v2, v3)
        for fi22 := 0; fi22 < 20; fi22++ {
            v2 += v2
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 int8 = -128
                v25 := []int{(v2 / (((fi22 ^ fi22)) | 1)), (fi23 * 2147483647), (fi23 ^ v2)}
                if true {
                    var v26 int64 = 0
                    _ = v26
                }
                switch -2147483648 {
                case -204:
                    fmt.Println(v24, v3)
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            v2 = (fi27 << ((255) & 63))
            if (1 != -343) {
                for fi28 := 0; fi28 < 4; fi28++ {
                    var v29 uint8 = 34
                    _ = len(v1)
                    _ = fi28
                    _ = v29
                }
                for fi30 := 0; fi30 < 12; fi30++ {
                    if (!(!((0 <= 42) && (0.1 >= 1.5)))) {
                        switch (("foo bar" + "a") + "") {
                        case "\n\t":
                            var v31 float64 = float64(-2.25)
                            _ = v31
                        case "foo bar":
                            v32 := GenS5{F0: true, F1: true, F2: -1}
                            _ = v32
                        }
                        fmt.Println(fi30, fi27)
                    } else {
                        v33 := GenS2{F0: 100, F1: 255}
                        var v34 uint64 = 62
                        v35 := v1["hello"]
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    if true {
                        fmt.Println(v2, fi27, fi30)
                        var v36 float64 = (((3.14159 + (float64((1.5 - 3.14159)) + 1e10)) - 1.5) / ((float64(3.14159) - (float64(1.5) / (1e10 * ((float64(1e10) * float64(1.5)) / 0.1)))) * 0.1))
                        v2 -= fi30
                        _ = v36
                    } else {
                        v3 = (v3 ^ (v3 << ((42) & 63)))
                        var v37 uint32 = (v3 >> ((255) & 63))
                        var v38 uint16 = 306
                        _ = v37
                        _ = v38
                    }
                    v3 = 0
                    _ = fi30
                }
            } else {
                v39 := []int{v2}
                v2 -= 32768
                _ = v39
            }
            if false {
                for fi40 := 0; fi40 < 10; fi40++ {
                    v41 := []int{fi27, 151, fi40}
                    var v42 int = 0
                    _ = len(v1)
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                v43 := v1["hello"]
                for fi44 := 0; fi44 < 1; fi44++ {
                    v3 = (v3 >> ((18446744073709551615) & 63))
                    v45 := GenS3{F0: v3}
                    _ = fi44
                    _ = v45
                }
                _ = v43
            } else {
                var v46 int64 = 1
                var v47 int16 = 32767
                v46 = (v46 - v46)
                _ = v46
                _ = v47
            }
            _ = fi27
        }
    } else {
        fmt.Println(v3, v2)
        if (1 >= 0) {
            v48 := GenS5{F0: true, F1: ("\n\t" > (("a" + "foo bar") + "")), F2: -1}
            var v49 int8 = 1
            _ = v48
            _ = v49
        }
        v2 -= v2
    }
    var v50 int16 = -47
    v51 := GenS3{F0: (^(v3))}
    fmt.Println(v2, v50)
    v1[""] = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v8
}
