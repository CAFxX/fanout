package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint; F1 string; F2 uint }
type GenS6 struct { F0 int }
func genFn0(p0 uint, p1 int) bool {
    var v6 float32 = 0.0
    v7 := GenS1{F0: (-10 >= 1000), F1: 0, F2: 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (true || true)
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    var v3 int64 = 32
    var v4 int32 = 2147483647
    var v5 uint64 = 991
    v5 += v5
    v8 := GenS1{F0: (-2147483648 < 1), F1: v4, F2: 4294967295}
    if v2 {
        var v9 uint8 = 42
        if (((1e10 / 1e10) == (((1.5 - ((0.0 + 3.14159) - 1.5)) + 0.0) * 1e10)) && (127 == 1)) {
            v4 += v4
        }
        v9 = 143
        _ = v9
    } else {
        v4 -= (v4 << ((525) & 63))
        v3 -= v3
        v5 += 133
    }
    if (614 > 2) {
        var v10 uint8 = 1
        _ = v10
    }
    fmt.Println(v2)
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := []int{fi11, (fi11 >> ((77) & 63))}
        if v2 {
            if v2 {
                v13 := v1["a"]
                _ = v13
            } else {
                v2 = v2
                switch (("hello" + "a") + "") {
                case "foo bar":
                    var v14 uint = 0
                    _ = v14
                case "\n\t":
                    if ((3.14159 + 1e10) == 0.1) {
                        fmt.Println(fi11, v5, v2)
                        fmt.Println(fi11)
                        var v15 uint = 42
                        _ = v15
                    }
                }
                v16 := v1["hello"]
                _ = v16
            }
            if ((!v2) && (1 > (v4 / (((v4 & (v4 % ((-2147483648) | 1)))) | 1)))) {
                if true {
                    delete(v1, "foo bar")
                    var v17 uint64 = 18446744073709551615
                    var v18 uint = 4294967295
                    _ = v17
                    _ = v18
                }
                if v2 {
                    v19 := v1[("" + ("" + "a"))]
                    v4 = (-(v4))
                    _ = v19
                } else {
                    if (2 >= 255) {
                        v2 = ("foo bar" == "a")
                    } else {
                        if ((v2 && (!(v2 && false))) || v2) {
                            v20 := GenS2{F0: (v5 - (v5 & (v5 + 4294967295))), F1: -9223372036854775808}
                            _ = v20
                        } else {
                            v3 = 781
                            var v21 uint16 = 32767
                            if (v4 == v4) {
                                _ = len(v1)
                                var v22 uint = 42
                                if (fi11 > 2147483647) {
                                    fmt.Println(v21)
                                } else {
                                    switch ((("foo bar" + "") + "a") + ("" + "foo bar")) {
                                    case "hello":
                                        v23 := GenS2{F0: (v5 ^ (^(v5))), F1: (-(fi11))}
                                        _ = v23
                                    }
                                }
                                _ = v22
                            } else {
                                v24 := GenS3{F0: 4, F1: 0, F2: (v5 & (^(v5)))}
                                v12 = append(v12, fi11)
                                v21 += (v21 - v21)
                                _ = v24
                            }
                            _ = v21
                        }
                    }
                    v2 = (true && (!(v5 >= v5)))
                    for fi25 := 0; fi25 < 15; fi25++ {
                        var v26 int8 = 124
                        _ = fi25
                        _ = v26
                    }
                }
            }
        } else {
            var v27 int = (fi11 - fi11)
            if v2 {
                v2 = (2147483647 >= v4)
                v28 := GenS4{F0: 32767, F1: true}
                var v29 int = v27
                _ = v28
                _ = v29
            } else {
                for fi30 := 0; fi30 < 17; fi30++ {
                    v31 := GenS5{F0: 255, F1: "\n\t", F2: 4294967295}
                    v32 := GenS3{F0: 65535, F1: 963, F2: 42}
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v27 = (^(fi11))
            }
            _ = v27
        }
        _ = fi11
        _ = v12
    }
    v3 -= -424
    v33 := GenS6{F0: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v8
}
