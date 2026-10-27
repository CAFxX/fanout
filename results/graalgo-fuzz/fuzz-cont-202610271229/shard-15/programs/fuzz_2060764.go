package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16, p1 int16, p2 uint32) int8 {
    var v4 uint8 = 1
    p2 = (p2 | (p2 << ((1) & 63)))
    p1 = 0
    for fi5 := 0; fi5 < 1; fi5++ {
        if (false || (-1 < 127)) {
            v4 -= (v4 << ((0) & 63))
            if false {
                p1 = (p1 / (((p1 / (((p1 + 1000)) | 1))) | 1))
                fmt.Println(p1, p0)
            } else {
                for fi6 := 0; fi6 < 1; fi6++ {
                    var v7 int64 = -1
                    var v8 string = "foo bar"
                    fmt.Println(v4, p1)
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                var v9 uint8 = (v4 + 41)
                v10 := []int{}
                _ = v10
                _ = v9
            }
            if (!(true || (v4 == (^(v4))))) {
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 uint32 = (^(p2))
                    if false {
                        if (("\n\t" + "a") == "a") {
                            v13 := GenS1{F0: -128}
                            var v14 float32 = 1.5
                            _ = v13
                            _ = v14
                        }
                    }
                    _ = fi11
                    _ = v12
                }
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS1{F0: 100}
                    p2 += 1
                    var v17 uint16 = p0
                    v17 += 0
                    v18 := GenS1{F0: 42}
                    v19 := []int{69}
                    var v20 int16 = p1
                    if true {
                        var v21 uint16 = 32767
                        v21 = (^(v21))
                        var v22 uint16 = 2
                        p1 = -709
                        v23 := GenS2{F0: -114}
                        p2 -= p2
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                var v24 string = "a"
                _ = v24
            } else {
                v25 := []int{(fi5 * (fi5 + fi5)), (^(fi5)), (fi5 | 32768), -859}
                _ = v25
            }
        } else {
            var v26 int32 = 0
            var v27 int32 = (v26 - v26)
            p1 += 32767
            _ = v26
            _ = v27
        }
        var v28 int16 = (p1 << ((42) & 63))
        v29 := []int{(fi5 % ((1) | 1)), (fi5 & (fi5 / (((^(fi5))) | 1))), fi5}
        _ = fi5
        _ = v28
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 37
}

func main() {
    var v1 uint = 1
    var v2 uint8 = 255
    var v3 bool = true
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := []int{987, fi30}
        var v32 float32 = 1.5
        _ = fi30
        _ = v31
        _ = v32
    }
    if v3 {
        var v33 int16 = 764
        var v34 int = 127
        _ = v33
        _ = v34
    } else {
        var v35 uint32 = 198
        if v3 {
            var v36 float32 = ((float32(-2.25) + (0.1 - -2.25)) - float32(float64(0.0)))
            v35 += v35
            _ = v36
        }
        _ = v35
    }
    fmt.Println(v3)
    var v37 bool = (((!(!false)) && v3) && (!false))
    for fi38 := 0; fi38 < 18; fi38++ {
        switch fi38 {
        case 1:
            var v39 string = "hello"
            _ = v39
        default:
            v3 = (-1 >= 0)
        }
        var v40 uint16 = 861
        v1 = 809
        _ = fi38
        _ = v40
    }
    if v3 {
        for fi41 := 0; fi41 < 15; fi41++ {
            for fi42 := 0; fi42 < 18; fi42++ {
                fmt.Println(v37)
                var v43 int = (fi41 ^ (fi42 >> (((v1 & v1)) & 63)))
                _ = fi42
                _ = v43
            }
            if v3 {
                v37 = false
                v44 := GenS2{F0: 823}
                _ = v44
            } else {
                fmt.Println(v1, v1, v2)
            }
            v2 += (v2 * 2)
            _ = fi41
        }
    } else {
        v45 := []int{}
        _ = v45
    }
    var v46 int = 9223372036854775807
    var v47 int16 = 856
    v48 := []int{}
    if true {
        v49 := GenS1{F0: 1}
        _ = v49
    } else {
        v3 = v3
        var v50 float32 = (0.0 + 3.14159)
        var v51 float64 = 0.0
        _ = v50
        _ = v51
    }
    var v52 bool = v3
    v37 = (v46 > (v46 / (((v46 % (((-(v46))) | 1))) | 1)))
    v2 += v2
    var v53 uint8 = v2
    if ((v46 + 912) < v46) {
        v46 += (v46 ^ v46)
        var v54 int32 = -1
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v46
    _ = v47
    _ = v48
    _ = v52
    _ = v53
}
