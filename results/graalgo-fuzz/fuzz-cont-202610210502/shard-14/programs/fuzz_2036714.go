package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int64 }
type GenS2 struct { F0 uint; F1 int; F2 bool }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 float32 = 1.5
    var v2 uint8 = 128
    var v3 int32 = -2147483648
    var v4 float32 = v1
    var v5 uint64 = 0
    v6 := []int{4294967295, -394}
    v7 := GenS1{F0: 0, F1: -2147483648, F2: -9223372036854775808}
    switch ("hello" + "a") {
    case "a":
        var v8 uint = 42
        _ = v8
    case "\n\t":
        v9 := GenS2{F0: 42, F1: 127, F2: ((1 == -1) && true)}
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 2; fi10++ {
            if len(v6) > 0 {
                v6[0] = fi10
            }
            _ = fi10
        }
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v4 = 1e10
        var v12 int8 = -1
        if false {
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS2{F0: 574, F1: 0, F2: (true && false)}
                if (uint16(v5) == 849) {
                    var v15 uint8 = v2
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if (!(0 >= (v5 / ((4294967295) | 1)))) {
                        if false {
                            var v16 int16 = -32768
                            _ = v16
                        } else {
                            v5 = v5
                            if (855 == 0) {
                                v4 = -2.25
                            } else {
                                var v17 uint16 = 32767
                                v18 := GenS1{F0: (v5 | v5), F1: 118, F2: 42}
                                _ = v17
                                _ = v18
                            }
                            if (true || (-19 == (v12 ^ (v12 ^ v12)))) {
                                var v19 uint8 = v2
                                fmt.Println(v12)
                                v12 += (^(v12))
                                _ = v19
                            }
                        }
                        v5 = (v5 >> ((48) & 63))
                        if true {
                            var v20 uint32 = 42
                            if (uint8(v5) > v2) {
                                var v21 float64 = 3.14159
                                var v22 int64 = 0
                                if true {
                                    if len(v6) > 0 {
                                        v6[0] = (^(fi13))
                                    }
                                    if (!(!true)) {
                                        v15 = (v15 & v15)
                                        var v23 float32 = v1
                                        if len(v6) > 0 {
                                            v6[0] = (fi13 | (fi13 & (fi13 / (((^(fi13))) | 1))))
                                        }
                                        var v24 int32 = v3
                                        var v25 bool = ((v5 * (v5 << ((244) & 63))) > (v5 >> ((65535) & 63)))
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    } else {
                                        switch -1 {
                                        case -142:
                                            fmt.Println(v21, v3)
                                        case 42:
                                            switch fi11 {
                                            case 42:
                                                v26 := GenS3{F0: 42, F1: int16(v22)}
                                                _ = v26
                                            case 1:
                                                fmt.Println(fi11)
                                            case -1:
                                                var v27 bool = false
                                                _ = v27
                                            }
                                        default:
                                            if ((!false) || false) {
                                                if true {
                                                    _ = len(v6)
                                                    var v28 int32 = v3
                                                    _ = v28
                                                } else {
                                                    var v29 uint16 = 65535
                                                    var v30 uint64 = v5
                                                    var v31 int8 = (v12 ^ (v12 | v12))
                                                    var v32 uint8 = v2
                                                    v21 = -2.25
                                                    var v33 uint32 = 65535
                                                    _ = v29
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                }
                                                v34 := GenS1{F0: 1, F1: fi13, F2: 1}
                                                _ = v34
                                            }
                                        }
                                        if (0.0 >= v21) {
                                            v21 = float64(v4)
                                            v3 -= v3
                                            v35 := GenS1{F0: v5, F1: (^(fi11)), F2: 9223372036854775807}
                                            _ = v35
                                        }
                                    }
                                    v4 -= v4
                                }
                                _ = v21
                                _ = v22
                            }
                            _ = v20
                        } else {
                            v36 := GenS2{F0: 255, F1: fi11, F2: true}
                            v37 := GenS4{F0: -556, F1: (v5 | v5)}
                            _ = v36
                            _ = v37
                        }
                    } else {
                        var v38 uint32 = uint32(v5)
                        if len(v6) > 0 {
                            v6[0] = (fi13 * 888)
                        }
                        var v39 bool = true
                        _ = v38
                        _ = v39
                    }
                    var v40 uint8 = uint8(v5)
                    if (((1e10 + 0.1) + 0.0) > (0.1 * 0.0)) {
                        if (!true) {
                            v15 = v40
                            var v41 uint = uint(v15)
                            fmt.Println(v41, v3)
                            _ = v41
                        } else {
                            v42 := GenS2{F0: 4294967295, F1: (fi11 >> ((416) & 63)), F2: ((v12 ^ (v12 << ((534) & 63))) == (v12 | 0))}
                            v43 := GenS1{F0: v5, F1: int(v3), F2: -9223372036854775808}
                            _ = v42
                            _ = v43
                        }
                    } else {
                        if false {
                            v3 = (v3 % (((v3 | v3)) | 1))
                            if false {
                                v1 -= float32(3.14159)
                                var v44 int = (fi13 >> ((500) & 63))
                                _ = v44
                            }
                            var v45 uint = 255
                            if false {
                                var v46 float32 = v1
                                var v47 float64 = 0.1
                                _ = v46
                                _ = v47
                            }
                            switch (fi11 + 127) {
                            case 127:
                                var v48 int16 = 32767
                                _ = v48
                            case -2147483648:
                                _ = len(v6)
                            default:
                                if len(v6) > 0 {
                                    v6[0] = 32768
                                }
                            }
                            var v49 int8 = v12
                            var v50 int32 = v3
                            switch (fi13 - (fi11 ^ int(v45))) {
                            case -588:
                                var v51 int32 = int32(v15)
                                _ = v51
                            }
                            _ = v45
                            _ = v49
                            _ = v50
                        } else {
                            v52 := GenS5{F0: 42}
                            fmt.Println(v5, v3)
                            _ = v52
                        }
                        v40 -= v15
                        var v53 int32 = v3
                        _ = v53
                    }
                    _ = v15
                    _ = v40
                } else {
                    v54 := GenS1{F0: (v5 & v5), F1: (fi11 - (^(fi11))), F2: 785}
                    var v55 uint64 = 161
                    v6 = append(v6, (fi11 | 4294967295))
                    if (-9223372036854775808 != 220) {
                        v12 -= v12
                    }
                    _ = v54
                    _ = v55
                }
                _ = fi13
                _ = v14
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            v4 = float32((1.5 / 3.14159))
        }
        _ = fi11
        _ = v12
    }
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    v56 := GenS3{F0: v3, F1: -1}
    v57 := GenS4{F0: 436, F1: (^(v5))}
    v3 = -1
    var v58 int32 = (^(v3))
    for fi59 := 0; fi59 < 4; fi59++ {
        v60 := GenS3{F0: v3, F1: 42}
        _ = fi59
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v6
    _ = v7
}
